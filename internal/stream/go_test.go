package stream_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.octolab.org/safe"

	. "go.octolab.org/toolset/testit/internal/stream"
)

func TestGoCompileProcess(t *testing.T) {
	tests := map[string]struct {
		input  string
		output string
	}{
		"compile": {
			input:  "testdata/compile/stdout.txt",
			output: "testdata/compile/stdout.golden",
		},
		"compile (cached)": {
			input:  "testdata/compile/stdout_cached.txt",
			output: "testdata/compile/stdout.golden",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			input, err := os.Open(test.input)
			require.NoError(t, err)
			defer safe.Close(input, func(err error) { require.NoError(t, err) })

			expected, err := ioutil.ReadFile(test.output)
			require.NoError(t, err)

			output := bytes.NewBuffer(nil)
			require.NoError(t, GoTestCompile(input, output).Process())

			assert.Equal(t, expected, output.Bytes())
		})
	}

	t.Run("bad writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		output := NewMockWriter(ctrl)
		output.EXPECT().Write(gomock.Any()).Return(0, errors.New("write to broken pipe"))

		input := strings.NewReader("go test output")
		require.Error(t, GoTestCompile(input, output).Process())
	})
}
