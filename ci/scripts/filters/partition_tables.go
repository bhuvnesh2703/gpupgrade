// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"regexp"
	"strings"
)

var WithClauseRegex = `(.*WITH\s\(tablename[^,]*,)(.*)`

func FormatWithClauseIfExisting(re *regexp.Regexp, line string) string {
	result := re.FindAllStringSubmatch(line, -1)
	if result == nil {
		return line
	}
	groups := result[0]
	// replace all occurrences of single quotes
	stringWithoutSingleQuotes := strings.ReplaceAll(groups[2], "'", "")

	return groups[1] + stringWithoutSingleQuotes
}
