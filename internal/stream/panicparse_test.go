package stream_test

import (
	"bytes"
	"flag"
	"io"
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
	f, err := os.Open("testdata/panic/stdout.txt")
	require.NoError(t, err)
	defer safe.Close(f, func(err error) { require.NoError(t, err) })

	var (
		in  io.Reader = f
		out           = bytes.NewBuffer(nil)
	)
	defer func() {
		f, err := os.OpenFile("testdata/panic/stdout.golden", os.O_RDWR, 0644)
		require.NoError(t, err)
		defer safe.Close(f, func(err error) { require.NoError(t, err) })

		if *update {
			_, err = f.Write(out.Bytes())
			require.NoError(t, err)
			return
		}

		expected, err := ioutil.ReadAll(f)
		require.NoError(t, err)
		assert.Equal(t, string(expected), out.String())
	}()

	require.NoError(t, GoTestStackTrace(in, out).Process())
}
