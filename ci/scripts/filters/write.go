// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"fmt"
	"io"
	"log"
)

func writeBufAndLine(out io.Writer, buf []string, line string) []string {
	// We want to keep this line. Flush and empty our buffer first.
	if len(buf) > 0 {
		write(out, buf...)
		buf = (buf)[:0]
	}

	write(out, line)

	return buf
}

func write(out io.Writer, lines ...string) {
	for _, line := range lines {
		_, err := fmt.Fprintln(out, line)
		if err != nil {
			log.Fatalf("writing output: %+v", err)
		}
	}
}
