package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecution(t *testing.T) {
	stderr, stdout = ioutil.Discard, ioutil.Discard

	t.Run("success", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 0, code) }
		os.Args = []string{"root", "version"}
		main()
	})

	t.Run("failure", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 1, code) }
		os.Args = []string{"root", "unknown"}
		main()
	})
}
