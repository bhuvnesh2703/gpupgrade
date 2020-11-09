// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func init() {
	replacementFuncs = []ReplacementFunc{
		ReplacePrecision,
	}
}

func Filter5x(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// there are lines in icw regression suite requiring buffer
	// to be atleast 10000000, so keeping it a little higher for now.
	scanner.Buffer(nil, 9800*4024)

	for scanner.Scan() {
		line := scanner.Text()

		for _, replacementFunc := range replacementFuncs {
			line = replacementFunc(line)
		}

		_, err := fmt.Fprintln(out, line)
		if err != nil {
			log.Fatalf("writing output: %v", err)
		}
	}

	if scanner.Err() != nil {
		log.Fatalf("scanning stdin: %+v", scanner.Err())
	}
}
