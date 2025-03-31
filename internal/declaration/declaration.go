package declaration

import (
	// standard
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	// 3rd-party
	"gopkg.in/yaml.v3"
)

type (
	Named map[string]string

	Declaration struct {
		Name        string   `json:"name" yaml:"name"`
		Package     string   `json:"package" yaml:"package"`
		Type        string   `json:"type" yaml:"type"`
		Values      []string `json:"values" yaml:"values"`
		NamedValues []Named  `json:"named-values" yaml:"named-values"`
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

	// --------------------------------------------------------------------------------------------
	if d.Name == "" {
		errs = append(errs, fmt.Errorf("'name' cannot be empty"))
	}
	if !isEnumNameValid(d.Name) {
		errs = append(errs, fmt.Errorf("'name' %q does not conform to naming conventions; must be an alphanumeric string starting with an uppercase letter", d.Name))
	}

	// --------------------------------------------------------------------------------------------
	if d.Package == "" {
		errs = append(errs, fmt.Errorf("'package' cannot be empty"))
	}
	if !isPackageNameValid(d.Package) {
		errs = append(errs, fmt.Errorf("'package' %q does not conform to naming conventions; must contain only lowercase letters and underscores", d.Package))
	}

	// --------------------------------------------------------------------------------------------
	normalizedType := strings.ToLower(d.Type)
	if normalizedType != "string" && normalizedType != "int" {
		if d.Type == "" {
			errs = append(errs, fmt.Errorf("'type' cannot be empty; must be 'string' or 'int'"))
		} else {
			errs = append(errs, fmt.Errorf("invalid 'type' %q; must be 'string' or 'int'", d.Type))
		}
	}

	// --------------------------------------------------------------------------------------------
	if len(d.Values) > 0 && len(d.NamedValues) > 0 {
		errs = append(errs, fmt.Errorf("must have either 'values' or 'named-values', not both"))
	}

	keys := d.getKeys()
	if len(keys) == 0 {
		errs = append(errs, fmt.Errorf("must have at least one value in 'values' or 'named-values'"))
	}
	for _, k := range keys {
		if !isEnumKeyValid(k) {
			errs = append(errs, fmt.Errorf("item key %q does not conform to naming conventions; must be an alphanumeric string starting with an uppercase letter", k))
		}
	}

	// Check named-values types if enum type is int
	if normalizedType == "int" && len(d.NamedValues) > 0 {
		for _, nv := range d.NamedValues {
			for key, valStr := range nv {
				if _, err := strconv.Atoi(valStr); err != nil {
					errs = append(errs, fmt.Errorf("named-value %q (%s) for type 'int' must be an integer, got %q", key, d.Name, valStr))
				}
			}
		}
	}

	// --------------------------------------------------------------------------------------------
	return len(errs) == 0, errs
}

func (d *Declaration) getKeys() []string {
	var keys []string

	if len(d.Values) > 0 {
		keys = append(keys, d.Values...)
	}

	if len(d.NamedValues) > 0 {
		for _, nv := range d.NamedValues {
			for k := range nv {
				keys = append(keys, k)
			}
		}
	}

	return keys
}
