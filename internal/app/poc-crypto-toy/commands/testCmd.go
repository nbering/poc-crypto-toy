package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type TestCmd struct {
	foo bool
}

func (*TestCmd) Name() string     { return "test" }
func (*TestCmd) Synopsis() string { return "Testing out the Google subcommands module." }
func (*TestCmd) Usage() string {
	return `test [-foo]:
	Doesn't really do anything, yet.
`
}

func (t *TestCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&t.foo, "foo", false, "foo the bar in the biz baz for the buzz")
}

func (t *TestCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("Foo!")
	fmt.Println()
	return subcommands.ExitSuccess
}
