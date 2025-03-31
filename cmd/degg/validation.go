package main

import (
	// standard
	"fmt"
	"path/filepath"
	"strings"

	// internal
	"github.com/upsetbit/degg/internal/declaration"
)

func lowercaseExt(p string) string {
	return strings.ToLower(filepath.Ext(p))
}

func identifyInputFormat(inputFile string) (declaration.Format, error) {
	ext := lowercaseExt(inputFile)

	switch ext {
	case ".json":
		return declaration.JSON, nil
	case ".yaml", ".yml":
		return declaration.YAML, nil
	}

	return declaration.UNKNOWN, fmt.Errorf("unsupported input format '%s'; must be '.json', '.yaml' or '.toml'", ext)
}

func checkOutputFile(outputFile string) error {
	if lowercaseExt(outputFile) != ".go" {
		return fmt.Errorf("output file must have '.go' extension")
	}

	return nil
}
