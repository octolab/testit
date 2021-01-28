package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "testit",
		Short: "extended `go test` for better experience",
		Long:  "Extended `go test` for better experience.",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}

	command.AddCommand(
		Golang(),
	)

	return &command
}
