package filters

import (
	"bytes"
	"testing"
)

func TestFilter5X(t *testing.T) {
	t.Run("it replaces double precision entries with XXXX", func(t *testing.T) {
		var in, out bytes.Buffer

		line := "90823904823 90.90 12.3023.232\n"
		in.WriteString(line)

		expected := "90823904823 90.XXXX 12.3023.232\n"
		Filter5x(&in, &out)
		if out.String() != expected {
			t.Errorf("wrote %q want %q", out.String(), expected)
		}
	})

	t.Run("it writes stdin to stdout", func(t *testing.T) {
		var in, out bytes.Buffer

		line := "hello\n"
		in.WriteString(line)

		expected := "hello\n"
		Filter5x(&in, &out)
		if out.String() != expected {
			t.Errorf("wrote %q want %q", out.String(), expected)
		}
	})
}
