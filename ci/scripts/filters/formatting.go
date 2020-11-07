// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"regexp"
	"strings"
)

var (
	formatters       []formatter
	lineRegexes      []*regexp.Regexp
	blockRegexes     []*regexp.Regexp
	replacementFuncs []ReplacementFunc
)

type ReplacementFunc func(line string) string

// function to identify if the line matches a pattern
type shouldFormatFunc func(line string) bool

// function to create a formatted string using the tokens
type formatFunc func(tokens []string) (string, error)

// identifier and corresponding formatting function
type formatter struct {
	shouldFormat shouldFormatFunc
	format       formatFunc
}

// hold the current tokens for the formatting function
type formatContext struct {
	tokens     []string
	formatFunc formatFunc
}

// is formatting currently in progress
func (f *formatContext) formatting() bool {
	return f.formatFunc != nil
}

func (f *formatContext) addTokens(line string) {
	f.tokens = append(f.tokens, strings.Fields(line)...)
}

func endFormatting(line string) bool {
	return strings.Contains(line, ";")
}

func (f *formatContext) format() (string, error) {
	return f.formatFunc(f.tokens)
}

func newFormattingContext() *formatContext {
	return &formatContext{}
}

func (f *formatContext) find(formatters []formatter, line string) {
	if f.formatFunc != nil {
		return
	}

	for _, x := range formatters {
		if x.shouldFormat(line) {
			f.formatFunc = x.format
			break
		}
	}
}
