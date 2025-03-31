package main

import (
	// standard
	"fmt"
	"go/format"
	"os"

	// internal
	"github.com/upsetbit/degg/cmd/degg/cli"
	"github.com/upsetbit/degg/internal/declaration"
	"github.com/upsetbit/degg/internal/generator"
	"github.com/upsetbit/degg/internal/system"
)

func main() {
	c := cli.New(mainWork)

	if err := c.Act(); err != nil {
		fmt.Printf("\nError: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("\nEnum generated and formatted successfully!")
}

func mainWork(inputFile string, outputFile string) error {
	if err := system.ResolvePath(&inputFile); err != nil {
		return fmt.Errorf("failed to resolve input file path: %w", err)
	}

	if err := system.ResolvePath(&outputFile); err != nil {
		return fmt.Errorf("failed to resolve output file path: %w", err)
	}

	formatInput, err := identifyInputFormat(inputFile)
	if err != nil {
		return err
	}

	if err = checkOutputFile(outputFile); err != nil {
		return err
	}

	decl, err := openAndValidateDeclaration(formatInput, inputFile)
	if err != nil {
		return fmt.Errorf("failed to process declaration file '%s'", inputFile)
	}

	generatedCode, err := generator.Run(decl)
	if err != nil {
		return fmt.Errorf("failed to generate enum code: %w", err)
	}

	formattedCode, err := format.Source([]byte(generatedCode))
	if err != nil {
		_ = os.WriteFile(outputFile, []byte(generatedCode), 0644)
		return fmt.Errorf("failed to format generated code (unformatted code written to %s): %w", outputFile, err)
	}

	if err := os.WriteFile(outputFile, formattedCode, 0644); err != nil {
		return fmt.Errorf("failed to write output file '%s': %w", outputFile, err)
	}

	fmt.Printf("Successfully generated and formatted enum '%s' to '%s'\n", decl.Name, outputFile)

	return nil
}

func openAndValidateDeclaration(format declaration.Format, inputFile string) (*declaration.Declaration, error) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file '%s': %w", inputFile, err)
	}

	decl, err := declaration.From(input, format)
	if err != nil {
		return nil, fmt.Errorf("failed to parse declaration: %w", err)
	}

	if declarationIsValid, errs := decl.Validate(); !declarationIsValid {
		fmt.Println("Declaration validation failed:")
		for _, validationErr := range errs {
			fmt.Printf("- %s\n", validationErr)
		}
		return nil, fmt.Errorf("declaration is invalid")
	}

	fmt.Printf("Declaration '%s' loaded and validated successfully.\n", decl.Name)
	return decl, nil
}
