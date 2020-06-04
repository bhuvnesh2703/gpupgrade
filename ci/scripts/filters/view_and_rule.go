// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0
package filters

import (
	"regexp"
	"strings"
)

var (
	// regex for views/rule transformation
	spaceBetweenOpenBracketAndCharRegex *regexp.Regexp
	spaceBetweenCloseBracketRegex       *regexp.Regexp
	ruleOrViewCommentRegex              *regexp.Regexp
	ruleOrViewCreateRegex               *regexp.Regexp
	viewReplacementRegex                []*replacer
)

type replacer struct {
	Regex       *regexp.Regexp
	Replacement string
}

func (t *replacer) replace(line string) string {
	return t.Regex.ReplaceAllString(line, t.Replacement)
}

// viewReplacementPatterns is a map of regex substitutions.
var viewReplacementPatterns = map[string]string{
	`'(LT|HT)'::text`:                             `'${1}'`,
	`d_xpect_setup_1\.`:                           `d_xpect_setup.`,
	`public.d_xpect_setup d_xpect_setup_1`:        `public.d_xpect_setup`,
	`testtable(\d+)_1\.`:                          `testtable${1}.`,
	`qp_misc_rio.testtable(\d+) testtable(\d+)_1`: `qp_misc_rio.testtable${1}`,
	`cte cte_1`:                                   `cte`,
	`cte_1\.`:                                     `cte.`,
}

func init() {
	spaceBetweenOpenBracketAndCharRegex = regexp.MustCompile(`\(\s`)
	spaceBetweenCloseBracketRegex = regexp.MustCompile(`\)\s\)`)
	ruleOrViewCommentRegex = regexp.MustCompile(`; Type: (VIEW|RULE);`)
	ruleOrViewCreateRegex = regexp.MustCompile(`CREATE (VIEW|RULE)`)

	for regex, replacement := range viewReplacementPatterns {
		viewReplacementRegex = append(viewReplacementRegex, &replacer{
			Regex:       regexp.MustCompile(regex),
			Replacement: replacement,
		})
	}
}

func StartFormattingViewOrRuleDdlStmtIfExisting(buf []string, line string) bool {
	return len(buf) > 0 && ruleOrViewCommentRegex.MatchString(strings.Join(buf[:], " ")) && ruleOrViewCreateRegex.MatchString(line)
}

func FormatViewOrRuleDdl(allTokens []string) string {
	var line string
	if allTokens[1] == "RULE" {
		line = strings.Join(allTokens[:], " ")
	} else {
		line = strings.Join(allTokens[:4], " ") + "\n" + strings.Join(allTokens[4:], " ")
		for _, r := range viewReplacementRegex {
			line = r.replace(line)
		}
	}
	line = spaceBetweenOpenBracketAndCharRegex.ReplaceAllString(line, `(`)
	line = spaceBetweenCloseBracketRegex.ReplaceAllString(line, `))`)
	return line
}

func BuildViewOrRuleDdl(line string, allTokens *[]string) (string, bool) {
	tokens := strings.Fields(line)
	*allTokens = append(*allTokens, tokens...)
	formattingFinished := false
	// if the DDL terminator exists process the whole DDL statement
	if strings.Contains(line, ";") {
		formattingFinished = true
		return FormatViewOrRuleDdl(*allTokens), formattingFinished
	}
	return "", formattingFinished
}
