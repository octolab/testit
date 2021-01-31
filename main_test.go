package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/safe"
	cli "go.octolab.org/toolkit/cli/errors"
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

	t.Run("shutdown silent", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 2, code) }
		safe.Do(func() error { return cli.Silent{Code: 2, Message: "silence"} }, shutdown)
	})

	t.Run("shutdown with panic", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 1, code) }
		safe.Do(func() error { panic("test") }, shutdown)
	})
}
