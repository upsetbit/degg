package declaration

import (
	// standard
	"encoding/json"
	"fmt"

	// 3rd-party
	"gopkg.in/yaml.v3"
)

type (
	named map[string]string

	Declaration struct {
		Name        string   `json:"name" yaml:"name"`
		Package     string   `json:"package" yaml:"package"`
		Type        string   `json:"type" yaml:"type"`
		Values      []string `json:"values" yaml:"values"`
		NamedValues []named  `json:"named-values" yaml:"named-values"`
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
		errs = append(errs, fmt.Errorf("'name' does not conform to naming conventions; must be an alphanumeric string starting with an uppercase letter"))
	}

	// --------------------------------------------------------------------------------------------
	if d.Package == "" {
		errs = append(errs, fmt.Errorf("'package' cannot be empty"))
	}
	if !isPackageNameValid(d.Package) {
		errs = append(errs, fmt.Errorf("'package' does not conform to naming conventions; must contain only lowercase letters and underscores"))
	}

	// --------------------------------------------------------------------------------------------
	if len(d.Values) > 0 && len(d.NamedValues) > 0 {
		errs = append(errs, fmt.Errorf("must have either values or named-values, not both"))
	}

	keys := d.getKeys()
	if len(keys) == 0 {
		errs = append(errs, fmt.Errorf("must have at least one value or named-value"))
	}
	for _, k := range keys {
		if !isEnumKeyValid(k) {
			errs = append(errs, fmt.Errorf("item '%s' does not conform to naming conventions; must be an alphanumeric string starting with an uppercase letter", k))
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
