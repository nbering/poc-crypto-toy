package commands

import (
	"context"
	"encoding/asn1"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

// ParseCmd is the subcommand implementation for debugging raw ANS.1 content.
type ParseCmd struct {
	in string
}

// Name provides the command name for the subcommand implementation.
func (*ParseCmd) Name() string { return "parse" }

// Synopsis provides a brief summary of the subcommands functionality for
// the CLI subcommand implementation.
func (*ParseCmd) Synopsis() string {
	return "Read an ASN.1 structure from a file and describe to stdout."
}

// Usage provides the subcommand help text for the CLI.
func (*ParseCmd) Usage() string {
	return `parse [-in]:
	Parse a file as an ASN.1 structure.
`
}

// SetFlags provides the command flags specific to the parse subcommand.
func (p *ParseCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.in, "in", "-", "Path to input file, or '-' to read from STDIN.")
}

// Execute provides the actual command functionality for the parse subcommand.
func (p *ParseCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var rawData []byte
	if p.in == "-" {
		rawData, _ = ioutil.ReadAll(os.Stdin)
	} else {
		inFile, _ := filepath.Abs(p.in)
		rawData, _ = ioutil.ReadFile(inFile)
	}

	var decoded asn1.RawValue

	asn1.Unmarshal(rawData, &decoded)

	printAsn1(decoded)
	fmt.Println()
	return subcommands.ExitSuccess
}

func printAsn1(d asn1.RawValue) {
	fmt.Printf("Class: %v, Tag: %v, Length: %v, (%v)\n", d.Class, d.Tag, len(d.Bytes), tagToName(d.Tag))
	if d.IsCompound {
		var c asn1.RawValue
		var rest []byte = d.Bytes
		for {
			rest, _ = asn1.Unmarshal(rest, &c)
			printAsn1(c)

			if len(rest) < 1 {
				break
			}
		}
	}
}

func tagToName(t int) string {
	switch t {
	case asn1.TagBoolean:
		return "Boolean"
	case asn1.TagInteger:
		return "Integer"
	case asn1.TagSequence:
		return "Sequence"
	default:
		return "UNKNOWN"
	}
}
