package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

// New returns the new root command.
func New(stderr, stdout io.Writer) *cobra.Command {
	command := cobra.Command{
		Use:   "testit",
		Short: "extended testing toolset",
		Long:  "Extended testing toolset.",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}

	command.SetErr(stderr)
	command.SetOut(stdout)
	command.AddCommand(
		Golang(),
	)

	return &command
}
