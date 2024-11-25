package system

import (
	"os"
	"path/filepath"
)

func ResolvePath(p *string) error {
	if filepath.IsAbs(*p) {
		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	*p, err = filepath.Abs(filepath.Join(wd, *p))
	if err != nil {
		return err
	}

	return nil
}
