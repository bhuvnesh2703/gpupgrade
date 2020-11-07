// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

/*
	The filter command massages the post-upgrade SQL dump by removing known
	differences. It does this with the following set of rules

	- Line rules are regular expressions that will cause any matching lines to
	be removed immediately.

	- Block rules are regular expressions that cause any matching lines, and any
	preceding comments or blank lines, to be removed.

	- Formatting rules are a set of functions that can format the sql statement tokens
	into a desired format

	filter reads from stdin and writes to stdout. Usage:

		filter < target.sql > target-filtered.sql

	Error handling is basic: any failures result in a log.Fatal() call.
*/
package main

import (
	"os"

	"github.com/greenplum-db/gpupgrade/ci/scripts/filters"
)

func main() {
	filters.Filter6x(os.Stdin, os.Stdout)
}
