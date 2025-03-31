package generator

import (
	"embed"
	"path"
)

var (
	//go:embed templates/*
	templates embed.FS
)

func readTemplate(filename string) (*string, error) {
	p := path.Join("templates", filename+".tmpl")
	b, err := templates.ReadFile(p)
	if err != nil {
		return nil, err
	}

	s := string(b)
	return &s, nil
}
