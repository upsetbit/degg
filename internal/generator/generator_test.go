package generator_test

import (
	"strings"
	"testing"

	"github.com/upsetbit/degg/internal/declaration"
	"github.com/upsetbit/degg/internal/generator"
)

func assertContains(t *testing.T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected string to contain %q, but it did not. String: %s", substr, s)
	}
}

func TestRun_StringEnum_Values(t *testing.T) {
	decl := &declaration.Declaration{
		Name:    "Color",
		Package: "colors",
		Type:    "string",
		Values:  []string{"RED", "GREEN", "BLUE"},
	}

	code, err := generator.Run(decl)
	if err != nil {
		t.Fatalf("generator.Run failed: %v", err)
	}

	assertContains(t, code, "package colors")
	assertContains(t, code, "\tColor string\n")
	assertContains(t, code, "RED Color = \"RED\"")
	assertContains(t, code, "GREEN Color = \"GREEN\"")
	assertContains(t, code, "BLUE Color = \"BLUE\"")
	assertContains(t, code, "ErrInvalidColor = errors.New(\"invalid value for Color, must be one of [RED, GREEN, BLUE]\")")
	assertContains(t, code, "func FromValue(c string) (Color, error)")
	assertContains(t, code, "case \"RED\":")
	assertContains(t, code, "func FromName(c string) (Color, error)")
	assertContains(t, code, "switch strings.ToUpper(c)")
	assertContains(t, code, "case \"GREEN\":")
	assertContains(t, code, "func (c Color) String() string")
	assertContains(t, code, "return string(c)")
	assertContains(t, code, "func (c Color) Int() int")
	assertContains(t, code, "case BLUE:")
	assertContains(t, code, "return 2")
}

func TestRun_IntEnum_NamedValues(t *testing.T) {
	decl := &declaration.Declaration{
		Name:    "Status",
		Package: "process",
		Type:    "int",
		NamedValues: []declaration.Named{
			{"PENDING": "10"},
			{"RUNNING": "20"},
			{"FAILED": "-1"},
		},
	}

	code, err := generator.Run(decl)
	if err != nil {
		t.Fatalf("generator.Run failed: %v", err)
	}

	assertContains(t, code, "package process")
	assertContains(t, code, "PENDING Status = 10")
	assertContains(t, code, "RUNNING Status = 20")
	assertContains(t, code, "FAILED Status = -1")
	assertContains(t, code, "ErrInvalidStatus = errors.New(\"invalid value for Status, must be one of [PENDING, RUNNING, FAILED]\")")
	assertContains(t, code, "func FromValue(s int) (Status, error)")
	assertContains(t, code, "case 10:")
	assertContains(t, code, "return PENDING, nil")
	assertContains(t, code, "func FromName(s string) (Status, error)")
	assertContains(t, code, "case \"FAILED\":")
	assertContains(t, code, "return FAILED, nil")
	assertContains(t, code, "func (s Status) String() string")
	assertContains(t, code, "case RUNNING:")
	assertContains(t, code, "return \"RUNNING\"")
	assertContains(t, code, "func (s Status) Int() int")
	assertContains(t, code, "return int(s)")
}

func TestRun_ErrorOnNilDeclaration(t *testing.T) {
	_, err := generator.Run(nil)
	if err == nil {
		t.Fatal("Expected an error when declaration is nil, but got nil")
	}

	expectedErrMsg := "declaration cannot be nil"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message to contain %q, got %q", expectedErrMsg, err.Error())
	}
}

func TestRun_ErrorOnMissingValues(t *testing.T) {
	decl := &declaration.Declaration{
		Name:    "Empty",
		Package: "test",
		Type:    "string",
	}

	_, err := generator.Run(decl)
	if err == nil {
		t.Fatal("Expected an error when declaration is missing values, but got nil")
	}
	expectedErrMsg := "declaration must contain 'values' or 'named-values'"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message to contain %q, got %q", expectedErrMsg, err.Error())
	}
}
