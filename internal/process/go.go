package process

import (
	"context"
	"io"
	"os"
	"os/exec"

	"go.octolab.org/unsafe"
)

type Option = func(cmd *exec.Cmd) error

func GoTestCompile(ctx context.Context, options ...Option) (*exec.Cmd, error) {
	cmd := exec.CommandContext(ctx, "go", "test", "-run", "^Fake$$")

	for _, configure := range options {
		if err := configure(cmd); err != nil {
			return nil, err
		}
	}

	return cmd, nil
}

func WithArgs(args []string) Option {
	return func(cmd *exec.Cmd) error {
		cmd.Args = append(cmd.Args, args...)
		return nil
	}
}

func WithCurrentEnv() Option {
	return func(cmd *exec.Cmd) error {
		cmd.Env = os.Environ()
		return nil
	}
}

func WithSignalPropagation(ctx context.Context, signals <-chan os.Signal) Option {
	return func(cmd *exec.Cmd) error {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case sig := <-signals:
					unsafe.Ignore(cmd.Process.Signal(sig))
				}
			}
		}()
		return nil
	}
}

func WithStderr(output io.Writer) Option {
	return func(cmd *exec.Cmd) error {
		cmd.Stderr = output
		return nil
	}
}

func WithStdin(input io.Reader) Option {
	return func(cmd *exec.Cmd) error {
		cmd.Stdin = input
		return nil
	}
}

func WithStdout(output io.Writer) Option {
	return func(cmd *exec.Cmd) error {
		cmd.Stdout = output
		return nil
	}
}
