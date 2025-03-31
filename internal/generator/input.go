package generator

import (
	"fmt"
	"strings"

	"github.com/upsetbit/degg/internal/declaration"
)

type (
	Value struct {
		Key   string
		Value string
		IsInt bool
	}

	GeneratorInput struct {
		Package        string
		EnumName       string
		EnumType       string // "string" or "int"
		EnumNameLetter string
		IsTypeString   bool
		IsTypeInt      bool
		Values         []Value
		ErrorInvalid   string // e.g., "InvalidColorErr"
		AllValues      string // e.g., "[RED, BLUE, GREEN]"
	}
)

func declarationToGeneratorInput(d *declaration.Declaration) (*GeneratorInput, error) {
	if d == nil {
		return nil, fmt.Errorf("declaration cannot be nil")
	}

	if len(d.Values) == 0 && len(d.NamedValues) == 0 {
		return nil, fmt.Errorf("declaration must contain 'values' or 'named-values'")
	}

	gi := initializeGeneratorInput(d)

	values, allKeys := processValues(d, gi.IsTypeInt, gi.IsTypeString)
	gi.Values = values
	gi.AllValues = formatAllValues(allKeys)

	return &gi, nil
}

func initializeGeneratorInput(d *declaration.Declaration) GeneratorInput {
	gi := GeneratorInput{
		Package:        d.Package,
		EnumName:       d.Name,
		EnumType:       strings.ToLower(d.Type),
		EnumNameLetter: strings.ToLower(d.Name[0:1]),
		ErrorInvalid:   fmt.Sprintf("ErrInvalid%s", d.Name),
	}

	gi.IsTypeString = gi.EnumType == "string"
	gi.IsTypeInt = gi.EnumType == "int"

	if !gi.IsTypeString && !gi.IsTypeInt {
		gi.EnumType = "string"
		gi.IsTypeString = true
		gi.IsTypeInt = false
	}

	return gi
}

func processValues(d *declaration.Declaration, isTypeInt, isTypeString bool) ([]Value, []string) {
	if len(d.Values) > 0 {
		return processIndexedValues(d.Values, isTypeInt)
	}

	return processNamedValues(d.NamedValues, isTypeInt, isTypeString)
}

func processIndexedValues(values []string, isTypeInt bool) ([]Value, []string) {
	result := make([]Value, 0, len(values))
	keys := make([]string, 0, len(values))

	for i, key := range values {
		var valStr string
		if isTypeInt {
			valStr = fmt.Sprintf("%d", i)
		} else {
			valStr = fmt.Sprintf("%q", key)
		}

		result = append(result, Value{Key: key, Value: valStr, IsInt: isTypeInt})
		keys = append(keys, key)
	}

	return result, keys
}

func processNamedValues(namedValues []declaration.Named, isTypeInt, isTypeString bool) ([]Value, []string) {
	totalCapacity := 0
	for _, m := range namedValues {
		totalCapacity += len(m)
	}

	result := make([]Value, 0, totalCapacity)
	keys := make([]string, 0, totalCapacity)

	for _, named := range namedValues {
		for key, valStr := range named {
			if isTypeString {
				valStr = fmt.Sprintf("%q", valStr)
			}

			result = append(result, Value{Key: key, Value: valStr, IsInt: isTypeInt})
			keys = append(keys, key)
		}
	}

	return result, keys
}

func formatAllValues(values []string) string {
	return "[" + strings.Join(values, ", ") + "]"
}
