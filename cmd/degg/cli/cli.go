package cli

import (
	// standard
	"os"

	// 3rd-party
	"github.com/urfave/cli/v2"
)

type (
	CLIActionCallback func(input string, output string) error

	CLI struct {
		app *cli.App
	}
)

const (
	flagInputFile  = "input"
	flagOutputFile = "output"
)

var (
	authors = []*cli.Author{
		{Name: "Caian Ertl", Email: "hi@caian.org"},
	}

	inputFlag = &cli.StringFlag{
		Name:     flagInputFile,
		Usage:    "the input file containing the enum definition",
		Aliases:  []string{"i"},
		Required: true,
	}

	outputFlag = &cli.StringFlag{
		Name:     flagOutputFile,
		Usage:    "the output file to write the generated code",
		Aliases:  []string{"o"},
		Required: true,
	}
)

func New(callback CLIActionCallback) *CLI {
	app := &cli.App{
		Name:     "degg",
		Usage:    "Dumb Enum Generator for Go",
		Version:  programVersion,
		Compiled: programCompiledAt,
		Authors:  authors,
		Flags:    []cli.Flag{inputFlag, outputFlag},

		Action: func(ctx *cli.Context) error {
			return callback(
				ctx.String(flagInputFile),
				ctx.String(flagOutputFile),
			)
		},
	}

	return &CLI{app}
}

func (c *CLI) Act() error {
	return c.app.Run(os.Args)
}
