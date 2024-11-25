package declaration

import (
	// standard
	"encoding/json"
	"fmt"

	// 3rd-party
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

type (
	named map[string]string

	Declaration struct {
		Name        string   `json:"name" yaml:"name" toml:"name"`
		Package     string   `json:"package" yaml:"package" toml:"package"`
		Type        string   `json:"type" yaml:"type" toml:"type"`
		Values      []string `json:"values" yaml:"values" toml:"values"`
		NamedValues named    `json:"named-values" yaml:"named-values" toml:"named-values"`
	}
)

func From(data []byte, format Format) (*Declaration, error) {
	var d Declaration
	var err error

	switch format {
	case JSON:
		err = json.Unmarshal(data, &d)
	case YAML:
		err = yaml.Unmarshal(data, &d)
	case TOML:
		err = toml.Unmarshal(data, &d)
	default:
		return nil, fmt.Errorf("unsupported format: %d", format)
	}

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (d *Declaration) Validate() (bool, []error) {
	var errs []error

	if d.Name == "" {
		errs = append(errs, fmt.Errorf("name is required"))
	}

	if d.Package == "" {
		errs = append(errs, fmt.Errorf("package is required"))
	}

	return len(errs) == 0, errs
}
