package cmd

import (
	"context"
	"io"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

		RunE: func(cmd *cobra.Command, args []string) error {
			job := new(async.Job)
			defer job.Wait()

			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			signals := make(chan os.Signal, 256)
			signal.Notify(signals)
			defer signal.Stop(signals)

			input, output := io.Pipe()
			defer safe.Close(output, unsafe.Ignore)

			valves := make([]stream.Valve, 0, 2)
			{
				set := pflag.NewFlagSet(cmd.Root().Use, pflag.ContinueOnError)
				colored := set.Bool("colored", false, "")
				stacked := set.Bool("stacked", false, "")
				unsafe.Ignore(set.Parse(args))
				args = exclude(args, "--colored", "--stacked")
				if *colored {
					valves = append(valves, stream.GoTest)
				}
				if *stacked {
					valves = append(valves, stream.GoTestStackTrace)
				}
			}

			task, err := process.GoTest(
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

			job.Do(
				stream.
					Connect(input, cmd.OutOrStdout()).
					Pipe(valves...).
					Operate,
				stream.Discard(input),
			)

			if err := task.Run(); err != nil {
				cmd.SilenceErrors = true
				return cli.Silent{Code: task.ProcessState.ExitCode()} // TODO:cli wrap error
			}
			return nil
		},
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

			signals := make(chan os.Signal, 256)
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

			job.Do(stream.GoTestCompile(input, cmd.OutOrStdout()).Operate, unsafe.Ignore)

			if err := task.Run(); err != nil {
				cmd.SilenceErrors = true
				return cli.Silent{Code: task.ProcessState.ExitCode()} // TODO:cli wrap error
			}
			return nil
		},
	}

	main.AddCommand(&compile)

	return &main
}

func exclude(input []string, by ...string) []string {
	filtered := input[:0]
	for _, str := range input {
		found := false
		for _, cmp := range by {
			if str == cmp {
				found = true
				break
			}
		}
		if !found {
			filtered = append(filtered, str)
		}
	}
	return filtered
}
