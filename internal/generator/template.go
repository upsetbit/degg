package generator

import (
	"bytes"
	"embed"
	"fmt"
	"path"
	"text/template"
)

var (
	//go:embed templates/*
	templates embed.FS
)

func processTemplate(templateFileName string, data *GeneratorInput) (string, error) {
	templatePath := path.Join("templates", templateFileName+".tmpl")

	t, err := template.ParseFS(templates, templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	return buf.String(), nil
}
