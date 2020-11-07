// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strings"
)

func init() {
	// linePatterns remove exactly what is matched, on a line-by-line basis.
	linePatterns := []string{
		`ALTER DATABASE .+ SET gp_use_legacy_hashops TO 'on';`,
		// TODO: There may be false positives because of the below
		// pattern, and we might have to do a look ahead to really identify
		// if it can be deleted.
		`START WITH \d`,
	}

	// blockPatterns remove lines that match, AND any comments or whitespace
	// immediately preceding them.
	blockPatterns := []string{
		"CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;",
		"COMMENT ON EXTENSION plpgsql IS",
		"COMMENT ON DATABASE postgres IS",
	}

	replacementFuncs = []ReplacementFunc{
		FormatWithClause,
		ReplacePrecision,
	}

	// patten matching functions and corresponding formatting functions
	formatters = []formatter{
		{shouldFormat: IsViewOrRuleDdl, format: FormatViewOrRuleDdl},
		{shouldFormat: IsTriggerDdl, format: FormatTriggerDdl},
	}

	for _, pattern := range linePatterns {
		lineRegexes = append(lineRegexes, regexp.MustCompile(pattern))
	}
	for _, pattern := range blockPatterns {
		blockRegexes = append(blockRegexes, regexp.MustCompile(pattern))
	}
}

func Filter6x(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// there are lines in icw regression suite requiring buffer
	// to be atleast 10000000, so keeping it a little higher for now.
	scanner.Buffer(nil, 9800*4024)

	var buf []string // lines buffered for look-ahead

	var formattingContext = newFormattingContext()

nextline:
	for scanner.Scan() {
		line := scanner.Text()

		formattingContext.find(formatters, line)
		if formattingContext.formatting() {
			formattingContext.addTokens(line)
			if endFormatting(line) {
				stmt, err := formattingContext.format()
				if err != nil {
					log.Fatalf("unexpected error: %#v", err)
				}
				buf = writeBufAndLine(out, buf, stmt)
				formattingContext = newFormattingContext()
			}
			continue nextline
		}

		// First filter on a line-by-line basis.
		for _, r := range lineRegexes {
			if r.MatchString(line) {
				continue nextline
			}
		}

		if strings.HasPrefix(line, "--") || len(line) == 0 {
			// A comment or an empty line. We only want to output this section
			// if the SQL it's attached to isn't filtered.
			buf = append(buf, line)
			continue nextline
		}

		for _, r := range blockRegexes {
			if r.MatchString(line) {
				// Discard this line and any buffered comment block.
				buf = buf[:0]
				continue nextline
			}
		}

		for _, replacementFunc := range replacementFuncs {
			line = replacementFunc(line)
		}

		buf = writeBufAndLine(out, buf, line)
	}

	if scanner.Err() != nil {
		log.Fatalf("scanning stdin: %+v", scanner.Err())
	}

	// Flush our buffer.
	if len(buf) > 0 {
		write(out, buf...)
	}
}
