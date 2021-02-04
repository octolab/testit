package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"go.octolab.org/safe"
	"go.octolab.org/toolkit/cli/cobra"
	"go.octolab.org/toolkit/cli/graceful"

	"go.octolab.org/toolset/testit/internal/cmd"
	"go.octolab.org/toolset/testit/internal/cnf"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
	exit    = os.Exit
	stderr  = io.Writer(os.Stderr)
	stdout  = io.Writer(os.Stdout)
)

func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	root := cmd.New(stderr, stdout)
	root.AddCommand(
		cobra.NewCompletionCommand(),
		cobra.NewVersionCommand(version, date, commit, cnf.Features...),
	)

	safe.Do(func() error { return root.ExecuteContext(ctx) }, graceful.Shutdown(stderr, exit))
}
