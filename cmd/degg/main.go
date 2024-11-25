package main

import (
	// standard
	"fmt"
	"os"

	// internal
	"github.com/upsetbit/degg/cmd/degg/cli"
	"github.com/upsetbit/degg/internal/declaration"
	"github.com/upsetbit/degg/internal/system"
)

func main() {
	c := cli.New(mainWork)

	if err := c.Act(); err != nil {
		fmt.Printf("\n%s\n", err.Error())
		os.Exit(1)
	}
}

func mainWork(input string, output string) error {
	if err := system.ResolvePath(&input); err != nil {
		return err
	}

	if err := system.ResolvePath(&output); err != nil {
		return err
	}

	format, err := identifyInputFormat(input)
	if err != nil {
		return err
	}

	if err = checkOutputFile(output); err != nil {
		return err
	}

	decl, err := openDeclaration(format, input)
	if err != nil {
		return nil
	}

	fmt.Printf("%v\n", decl)

	return nil
}

func openDeclaration(format declaration.Format, inputFile string) (*declaration.Declaration, error) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	decl, err := declaration.From(input, format)
	if err != nil {
		return nil, err
	}

	if declarationIsValid, errs := decl.Validate(); !declarationIsValid {
		for _, err := range errs {
			fmt.Println(err)
		}
		return nil, fmt.Errorf("declaration is invalid")
	}

	return decl, nil
}
