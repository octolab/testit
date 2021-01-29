package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "testit",
		Short: "extended testing toolset",
		Long:  "Extended testing toolset.",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}

	command.AddCommand(
		Golang(),
	)

	return &command
}
