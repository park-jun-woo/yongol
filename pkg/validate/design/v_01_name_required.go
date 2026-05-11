//ff:func feature=validate type=rule control=sequence topic=design-structural
//ff:what V-01 — name 필드 필수 검증
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// v01NameRequired validates that the design spec has a non-empty name field.
func v01NameRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.DesignSpec.Name != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    fs.DesignSpec.File,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[V-01] name field is required in DESIGN.md",
		Advice:  "Add a non-empty 'name' field to the YAML front matter",
	}}
}
