package declaration

import (
	"regexp"
)

var (
	enumNameRegex    = regexp.MustCompile(`^[A-Z][A-Za-z0-9_]*$`)
	packageNameRegex = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
)
