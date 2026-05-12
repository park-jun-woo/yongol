//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-structural
//ff:what V-04 — rounded, spacing 값이 유효한 dimension (px/em/rem) 또는 숫자 검증
package design

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var dimensionRe = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?(px|em|rem)?$`)

// v04DimensionValid validates that rounded and spacing values are valid dimensions.
func v04DimensionValid(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for name, val := range fs.DesignSpec.Rounded {
		if !dimensionRe.MatchString(val) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-04] rounded \"" + name + "\" has invalid dimension: " + val,
				Advice:  "Use a number optionally followed by px, em, or rem",
			})
		}
	}
	for name, val := range fs.DesignSpec.Spacing {
		if !dimensionRe.MatchString(val) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-04] spacing \"" + name + "\" has invalid dimension: " + val,
				Advice:  "Use a number optionally followed by px, em, or rem",
			})
		}
	}
	return diags
}
