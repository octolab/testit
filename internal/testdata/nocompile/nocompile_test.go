package nocompile_test

import (
	"math/rand"
	. "testdata/nocompile"
	"testing"
)

func Test(t *testing.T) {
	var id Interface

	id = rand.New(rand.NewSource(0))
	if !id.Unique() {
		t.FailNow()
	}
}
