// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

/*
	The filter command massages the post-upgrade SQL dump by removing known
	differences. Different set of rules are applied for dump from greenplum
	version 5 and 6. In general, the below set of rules are applied on the dump.

	- Line rules are regular expressions that will cause any matching lines to
	be removed immediately.

	- Block rules are regular expressions that cause any matching lines, and any
	preceding comments or blank lines, to be removed.

	- Formatting rules are a set of functions that can format the sql statement tokens
	into a desired format

	filter reads from an input file and writes to stdout. Usage:

		filter -version=5 -inputFile=dump.sql > dump-filtered.sql

	Error handling is basic: any failures result in a log.Fatal() call.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/greenplum-db/gpupgrade/ci/scripts/filters"
)

var (
	version6 = 6
	version5 = 5
	argCount = 2
)

type filter struct {
	filterFunc func(in io.Reader, out io.Writer)
}

func newFilter() *filter {
	return &filter{}
}

func (f *filter) setFunc(version int) {
	if version == version5 {
		f.filterFunc = filters.Filter5x
	} else {
		f.filterFunc = filters.Filter6x
	}
}

func main() {
	var (
		version   int
		inputFile string
	)

	flag.IntVar(&version, "version", 0, "identifier specific version of greenplum dump, i.e 5 or 6")
	flag.StringVar(&inputFile, "inputFile", "", "fully qualified input file name containing the dump")
	flag.Parse()

	if flag.NFlag() != argCount {
		fmt.Printf("requires %d arguments, got %d\n", argCount, flag.NFlag())
		flag.Usage()
		os.Exit(1)
	}

	if version != version5 && version != version6 {
		fmt.Printf("permitted -version values are %d and %d. but got %d\n", version5, version6, version)
		os.Exit(1)
	}

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Print(fmt.Errorf("%s: %w\n", inputFile, err))
		os.Exit(1)
	}

	filter := newFilter()
	filter.setFunc(version)
	filter.filterFunc(in, os.Stdout)
}
