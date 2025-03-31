package generator

import (
	"strings"
)

type (
	GeneratorInput struct {
		Package          string
		EnumName         string
		EnumNameLetter   string
		EnumKeysListed   string
		EnumValuesListed string
	}
)

func declarationToGeneratorInput(d *Decl) *GeneratorInput {
	gi := GeneratorInput{
		Package:        d.Package,
		EnumName:       d.Name,
		EnumNameLetter: strings.ToLower(d.Name[0:1]),
	}

	return &gi
}
