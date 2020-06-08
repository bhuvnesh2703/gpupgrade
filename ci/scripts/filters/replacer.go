// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import "regexp"

type Replacer struct {
	Regex           *regexp.Regexp
	ReplacementFunc func(re *regexp.Regexp, line string) string
}

func (t *Replacer) Replace(line string) string {
	return t.ReplacementFunc(t.Regex, line)
}
