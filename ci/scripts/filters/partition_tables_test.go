// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"regexp"
	"testing"
)

func Test_FormatWithClauseIfExisting(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "single quotes are removed and with clause ends without a comma",
			input:  "START ('2005-12-01 00:00:00'::timestamp without time zone) END ('2006-01-01 00:00:00'::timestamp without time zone) EVERY ('1 mon'::interval) WITH (tablename='order_lineitems_1_prt_2', appendonly='true', compresstype=quicklz, orientation='column' )",
			result: "START ('2005-12-01 00:00:00'::timestamp without time zone) END ('2006-01-01 00:00:00'::timestamp without time zone) EVERY ('1 mon'::interval) WITH (tablename='order_lineitems_1_prt_2', appendonly=true, compresstype=quicklz, orientation=column )",
		},
		{
			name:   "single quotes are removed and with clause ends with a comma",
			input:  "START ('2005-12-01 00:00:00'::timestamp without time zone) END ('2006-01-01 00:00:00'::timestamp without time zone) EVERY ('1 mon'::interval) WITH (tablename='order_lineitems_1_prt_2', appendonly='true', compresstype=quicklz, orientation='column' ),",
			result: "START ('2005-12-01 00:00:00'::timestamp without time zone) END ('2006-01-01 00:00:00'::timestamp without time zone) EVERY ('1 mon'::interval) WITH (tablename='order_lineitems_1_prt_2', appendonly=true, compresstype=quicklz, orientation=column ),",
		},
	}

	re := regexp.MustCompile(WithClauseRegex)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatWithClauseIfExisting(re, tt.input); got != tt.result {
				t.Errorf("got %v, want %v", got, tt.result)
			}
		})
	}
}
