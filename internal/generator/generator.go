package generator

import (
	"fmt"

	"github.com/upsetbit/degg/internal/declaration"
)

type (
	Decl declaration.Declaration
)

func Run(d *declaration.Declaration) (string, error) {
	input, err := declarationToGeneratorInput(d)
	if err != nil {
		return "", fmt.Errorf("failed to prepare generator input: %w", err)
	}

	generatedCode, err := processTemplate("enum.go", input)
	if err != nil {
		return "", fmt.Errorf("failed to generate code from template: %w", err)
	}

	return generatedCode, nil
}
