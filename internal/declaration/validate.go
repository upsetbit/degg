package declaration

import (
	"regexp"
)

var (
	enumNameRegex    = regexp.MustCompile(`^[A-Z][A-Za-z0-9_]*$`)
	enumKeyRegex     = regexp.MustCompile(`^[A-Z][A-Za-z0-9_]*$`)
	packageNameRegex = regexp.MustCompile(`^[a-z][a-z_]*$`)
)

func isEnumNameValid(name string) bool {
	return enumNameRegex.MatchString(name)
}

func isEnumKeyValid(key string) bool {
	return enumKeyRegex.MatchString(key)
}

func isPackageNameValid(name string) bool {
	return packageNameRegex.MatchString(name)
}
