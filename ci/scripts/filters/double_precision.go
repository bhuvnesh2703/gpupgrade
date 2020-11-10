// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	comma                 = ","
	tab                   = "\t"
	precisionRegex        = regexp.MustCompile(`\.\d+`)
	castToDoubleRegex     = regexp.MustCompile(`\.\d+::double precision`)
	bracketsRegex         = regexp.MustCompile(`(.*?)({|\(+)(.*?)(}|\)+)(.*)`)
	excludeRegex          = regexp.MustCompile(`VALUES.* WITH \(tablename|perform pg_sleep|time.sleep`)
	replacePrecisionFuncs = []func(line string) string{ReplacePrecisionInTabDelimitedData, ReplacePrecisionInBrackets}
)

func isFloat(n string) bool {
	if castToDoubleRegex.MatchString(n) {
		return true
	}

	_, err := strconv.ParseFloat(strings.TrimSpace(n), 64)
	return err == nil
}

// if the line contains certain patterns then don't replace the precision values,
// otherwise try to find the groups based on input regexp
func matchingGroups(r *regexp.Regexp, line string) [][]string {
	if excludeRegex.MatchString(line) {
		return nil
	}

	groups := r.FindAllStringSubmatch(line, -1)
	if groups == nil {
		return nil
	}

	return groups
}

// replaces precision will iterate over all the input tokens,
// and replaces all the tokens which are double precision floating
// point numbers with XX in place of precision values
func replacePrecisions(tokens []string, delimiter string) string {
	var formattedTokens []string
	for _, t := range tokens {
		if isFloat(t) {
			formattedTokens = append(formattedTokens, precisionRegex.ReplaceAllString(t, ".XX"))
		} else {
			formattedTokens = append(formattedTokens, t)
		}
	}

	return strings.Join(formattedTokens, delimiter)
}

// replace precision values in tab separated data
// ex:
//	input: 	122	12.22	abc
// 	output:	122	12.XX	abc
func ReplacePrecisionInTabDelimitedData(line string) string {
	tokens := strings.Split(line, tab)
	return replacePrecisions(tokens, tab)
}

// for a line [(-122.1204,37.267000000000003),(-122.123,37.271000000000001)],
// the below groups will be found:
// 	group[0][1] = [
// 	group[0][2] = (
// 	group[0][3] = -122.1204,37.267000000000003
// 	group[0][4] = )
// 	group[0][5] = ,(-122.123,37.271000000000001)]
// function ReplacePrecisionInBrackets captures all the groups, recursively
// processes group[0][5], where group[0][3] will contain the values which
// can be  in scope for replacement of precision. It then outputs a
// newly formatted string with precision values replaced with .XX
func ReplacePrecisionInBrackets(line string) string {
	groups := matchingGroups(bracketsRegex, line)
	if groups == nil {
		return line
	}

	return groups[0][1] +
		groups[0][2] +
		replacePrecisions(strings.Split(groups[0][3], comma), comma) +
		groups[0][4] +
		ReplacePrecisionInBrackets(groups[0][5]) // recursive call
}

func ReplacePrecision(line string) string {
	for _, f := range replacePrecisionFuncs {
		line = f(line)
	}

	return line
}
