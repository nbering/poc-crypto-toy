package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/nbering/poc-crypto-toy/internal/app/poc-crypto-toy/commands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.TestCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
