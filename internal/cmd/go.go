package cmd

import (
	"context"
	"errors"
	"io"
	"os"
	"os/signal"

	"github.com/maruel/panicparse/v2/stack"
	"github.com/spf13/cobra"
	"go.octolab.org/async"
	"go.octolab.org/safe"
	"go.octolab.org/strings"
	cli "go.octolab.org/toolkit/cli/cobra"
	"go.octolab.org/unsafe"

	"go.octolab.org/toolset/testit/internal/process"
	"go.octolab.org/toolset/testit/internal/stream"
)

func Golang() *cobra.Command {
	var (
		abspath bool
		colored bool
		stacked bool
	)

	test := cobra.Command{
		Use:   "go",
		Short: "proxy for go test with extra features",
		Long:  "Proxy for go test with extra features.",

		SilenceErrors:      false,
		DisableFlagParsing: true,
	}

	set := test.Flags()
	set.BoolVar(&abspath, "abspath", false, "replace relative paths by absolute")
	set.BoolVar(&colored, "colored", false, "enable colors")
	set.BoolVar(&stacked, "stacked", false, "enable panic and data race parsing")

	test.PreRunE = func(cmd *cobra.Command, args []string) error {
		if strings.PresentAny(args, "-h", "--help") {
			unsafe.Ignore(test.Usage())
			test.SilenceErrors = true
			return errors.New("shows usage")
		}
		return nil
	}

	test.RunE = func(cmd *cobra.Command, args []string) error {
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
			unsafe.Ignore(set.Parse(args))
			args = strings.Exclude(args, "--abspath", "--colored", "--stacked", "-help", "--help")
			if abspath {
				// TODO wrap stderr
			}
			if colored {
				valves = append(valves, stream.GoTest)
			}
			if stacked {
				valves = append(valves, stream.GoTestStackTrace(stack.DefaultOpts(), colored))
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

		return cli.SilentError(cmd, task.Run(), task.ProcessState.ExitCode())
	}

	test.AddCommand(compile())

	return &test
}

func compile() *cobra.Command {
	command := cobra.Command{
		Use:   "compile",
		Short: "make sure that all code is compiled",
		Long:  "Make sure that all code is compiled.",

		SilenceErrors:      false,
		DisableFlagParsing: true,
	}

	command.PreRunE = func(cmd *cobra.Command, args []string) error {
		if strings.PresentAny(args, "-h", "--help") {
			unsafe.Ignore(command.Usage())
			command.SilenceErrors = true
			return errors.New("shows usage")
		}
		return nil
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
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

		return cli.SilentError(cmd, task.Run(), task.ProcessState.ExitCode())
	}

	return &command
}
