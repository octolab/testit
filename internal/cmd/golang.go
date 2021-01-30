package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
	"go.octolab.org/safe"

	"go.octolab.org/toolset/testit/internal/stream"
)

func Golang() *cobra.Command {
	main := cobra.Command{
		Use:   "go",
		Short: "proxy for go test with extra features",
		Long:  "Proxy for go test with extra features.",

		Args: cobra.NoArgs,
	}

	compile := cobra.Command{
		Use:   "compile",
		Short: "make sure that all code is compiled",
		Long:  "Make sure that all code is compiled.",

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				args = append(args, "./...")
			}
			bin := exec.CommandContext(cmd.Context(), "go", append([]string{"test", "-run", "^Fake$$"}, args...)...)
			bin.Stderr = cmd.ErrOrStderr()
			reader, err := bin.StdoutPipe()
			if err != nil {
				return err
			}

			go safe.Do(bin.Run, func(err error) {})

			return stream.GoCompileProcess(reader, cmd.OutOrStdout())
		},
	}

	main.AddCommand(&compile)

	return &main
}
