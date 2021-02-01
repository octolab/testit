package stream_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.octolab.org/safe"

	. "go.octolab.org/toolset/testit/internal/stream"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestGoTestStackTrace(t *testing.T) {
	t.SkipNow()

	input, err := os.Open("testdata/panic/stdout.txt")
	require.NoError(t, err)
	defer safe.Close(input, func(err error) { require.NoError(t, err) })

	output := bytes.NewBuffer(nil)
	defer func() {
		golden, err := os.OpenFile("testdata/panic/stdout.golden", os.O_RDWR, 0644)
		require.NoError(t, err)
		defer safe.Close(golden, func(err error) { require.NoError(t, err) })

		if *update {
			_, err = golden.Write(output.Bytes())
			require.NoError(t, err)
			return
		}

		expected, err := ioutil.ReadAll(golden)
		require.NoError(t, err)
		assert.Equal(t, string(expected), output.String())
	}()

	require.NoError(t, GoTestStackTrace(input, output).Process())
}
