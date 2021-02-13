package failed_test

import (
	. "testdata/failed"
	"testing"
)

func TestConcat(t *testing.T) {
	tests := map[string]struct {
		input  []string
		output string
	}{
		"valid":   {[]string{"a", "b", "c"}, "abc"},
		"invalid": {[]string{"c", "b", "a"}, "abc"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if output := Concat(test.input...); test.output != output {
				t.Errorf("%q != %q", test.output, output)
			}
		})
	}
}
