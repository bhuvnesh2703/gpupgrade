package filters

import "testing"

func TestReplacePrecision(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "replaces occurrence of double precision",
			input:    "0000 abd 8902342 92342.2342 127.0.0.1",
			expected: "0000 abd 8902342 92342.XXXX 127.0.0.1",
		},
	}

	for _, c := range cases {
		output := ReplacePrecision(c.input)
		if output != c.expected {
			t.Errorf("got %q, want %q", output, c.expected)
		}
	}
}
