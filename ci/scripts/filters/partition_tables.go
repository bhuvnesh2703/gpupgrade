// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0
package filters

import "regexp"

var (
	// regex for partition table with clause
	singleQuoteRegex              *regexp.Regexp
	partitionTableWithClauseRegex *regexp.Regexp
)

func init() {
	singleQuoteRegex = regexp.MustCompile("'")
	partitionTableWithClauseRegex = regexp.MustCompile(`(.+)WITH \(tablename='(.[^,]*?)', (.*) \)(\,*)`)
}

func FormatWithClauseIfExisting(line string) string {
	result := partitionTableWithClauseRegex.FindAllStringSubmatch(line, -1)
	if result == nil {
		return line
	}
	groups := result[0]
	stringWithTableReplacement := "WITH (tablename='" + groups[2] + "'"
	// replace all occurrences of single quotes
	stringWithoutSingleQuotes := singleQuoteRegex.ReplaceAllString(groups[3], "")

	return groups[1] + stringWithTableReplacement + ", " + stringWithoutSingleQuotes + " )" + groups[4]
}
