package cli

import (
	// standard
	"fmt"
	"time"
)

var (
	programTag = "0.0.1"

	// compile-time
	ProgramCommitSHA = ""
	ProgramBuildTime = ""

	// run-time
	programVersion    = parseVersionOrDie(programTag, ProgramCommitSHA)
	programCompiledAt = parseCompiledAtOrDie(ProgramBuildTime)
)

func parseCompiledAtOrDie(ts string) time.Time {
	if dt, err := time.Parse("2006-01-02T15:04:05", ts); err != nil {
		panic(err)
	} else {
		return dt
	}
}

func parseVersionOrDie(tag string, sha string) string {
	if len(sha) == 0 {
		panic("missing commit SHA")
	}

	return fmt.Sprintf("%s (%s)", tag, sha)
}
