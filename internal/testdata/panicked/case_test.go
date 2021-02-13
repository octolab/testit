package panicked_test

import (
	. "testdata/panicked"
	"testing"
)

func TestTheGood_Divide(t *testing.T) {
	_, err := TheGood{}.Divide(5, 0)
	if err != nil {
		t.FailNow()
	}
}

func TestTheBad_Divide(t *testing.T) {
	TheBad{}.Divide(5, 0)
}

func TestTheUgly_Divide(t *testing.T) {
	TheUgly{}.Divide(5, 0)
}
