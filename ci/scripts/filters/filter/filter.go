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
	"flag"
	"fmt"
	"os"

	"github.com/greenplum-db/gpupgrade/ci/scripts/filters"
)

var (
	version6 = 6
	argCount = 2
)

func main() {
	var (
		version   int
		inputFile string
	)

	flag.IntVar(&version, "version", 0, "input file contains dump of greenplum version 5 or 6")
	flag.StringVar(&inputFile, "inputFile", "", "fully qualified input file name containing the dump")
	flag.Parse()

	if flag.NFlag() != argCount {
		fmt.Printf("requires %d arguments, got %d\n", argCount, flag.NFlag())
		flag.Usage()
		os.Exit(1)
	}

	if version != version6 {
		fmt.Printf("permitted -version values is: %d. but got %d", version6, version)
		os.Exit(1)
	}

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Print(fmt.Errorf("%s: %w", inputFile, err))
		os.Exit(1)
	}

	filters.Filter6x(in, os.Stdout)
}
