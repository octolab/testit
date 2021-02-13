package notests_test

import (
	. "testdata/notests"
	"testing"
)

func TestSum(t *testing.T) {
	tests := map[string]struct {
		a, b     int
		expected int
	}{
		"positive": {1, 2, 3},
		"negative": {-1, -2, -3},
		"mixed":    {1, -2, -1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if obtained := Sum(test.a, test.b); test.expected != obtained {
				t.FailNow()
			}
		})
	}
}
