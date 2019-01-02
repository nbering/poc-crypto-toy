package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

// TestCmd is a subcommand implementation that just demonstrates subcommand usage.
type TestCmd struct {
	foo bool
}

// Name returns the name of the TestCmd subcommand used in the CLI implementation.
func (*TestCmd) Name() string { return "test" }

// Synopsis returns a short description of TestCmd used in the CLI implementation.
func (*TestCmd) Synopsis() string { return "Testing out the Google subcommands module." }

// Usage returns a brief usage description of TestCmd for the help text
// in the CLI implementation.
func (*TestCmd) Usage() string {
	return `test [-foo]:
	Doesn't really do anything, yet.
`
}

// SetFlags sets up the command line flags for the TestCmd subcommand.
func (t *TestCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&t.foo, "foo", false, "foo the bar in the biz baz for the buzz")
}

// Execute provides the actual command implementation of TestCmd.
func (t *TestCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("Foo!")
	fmt.Println()
	return subcommands.ExitSuccess
}
