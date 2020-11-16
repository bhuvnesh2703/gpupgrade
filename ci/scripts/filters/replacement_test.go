package filters

import (
	"testing"
)

func TestReplaceBitTypeInDDL(t *testing.T) {
	tests := []struct {
		name string
		line string
		want string
	}{
		{
			name: `append B to the pattern 'n'::"bit"`,
			line: `a39 bit(1) DEFAULT '0'::"bit" ENCODING`,
			want: `a39 bit(1) DEFAULT B'0'::"bit" ENCODING`,
		},
		{
			name: `append B to the pattern ('n'::"bit")`,
			line: `a40 bit varying(5) DEFAULT ('1'::"bit")::bit varying`,
			want: `a40 bit varying(5) DEFAULT (B'1'::"bit")::bit varying`,
		},
		{
			name: `does not append B to the pattern B'n'::"bit"`,
			line: `a39 bit(1) DEFAULT B'0'::"bit" ENCODING`,
			want: `a39 bit(1) DEFAULT B'0'::"bit" ENCODING`,
		},
		{
			name: `replaces create view statement`,
			line: `CREATE VIEW test_invalid_schema_creation_tab_v1 AS`,
			want: `		CREATE VIEW test_invalid_schema_creation_tab_v1 AS SELECT f1 FROM test_invalid_schema_creation_tab;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Replacements(tt.line)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
