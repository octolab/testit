package cmd

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.octolab.org/async"
	"go.octolab.org/safe"
	cli "go.octolab.org/toolkit/cli/errors"
	"go.octolab.org/unsafe"

	"go.octolab.org/toolset/testit/internal/process"
	"go.octolab.org/toolset/testit/internal/stream"
)

func Golang() *cobra.Command {
	main := cobra.Command{
		Use:   "go",
		Short: "proxy for go test with extra features",
		Long:  "Proxy for go test with extra features.",

		DisableFlagParsing: true,
	}

	compile := cobra.Command{
		Use:   "compile",
		Short: "make sure that all code is compiled",
		Long:  "Make sure that all code is compiled.",

		SilenceErrors:      false,
		DisableFlagParsing: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			job := new(async.Job)
			defer job.Wait()

			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			signals := make(chan os.Signal)
			signal.Notify(signals)
			defer signal.Stop(signals)

			input, output := io.Pipe()
			defer safe.Close(output, unsafe.Ignore)

			task, err := process.GoTestCompile(
				ctx,
				process.WithArgs(args),
				process.WithCurrentEnv(),
				process.WithSignalPropagation(ctx, signals),
				process.WithStdin(cmd.InOrStdin()),
				process.WithStderr(cmd.ErrOrStderr()),
				process.WithStdout(output),
			)
			if err != nil {
				return err
			}

			job.Do(stream.GoTestCompile(input, cmd.OutOrStdout()).Process, unsafe.Ignore)

			if err := task.Run(); err != nil {
				if ws, ok := task.ProcessState.Sys().(syscall.WaitStatus); ok {
					err = cli.Silent{Code: ws.ExitStatus()}
					cmd.SilenceErrors = true
				}
				return err
			}
			return nil
		},
	}

	main.AddCommand(&compile)

	return &main
}
