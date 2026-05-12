//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-structural
//ff:what V-02 — colors 값이 유효한 hex (#XXX / #XXXX / #XXXXXX / #XXXXXXXX) 검증
package design

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var hexColorRe = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)

// v02HexValid validates that all color values are valid hex codes.
func v02HexValid(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for name, val := range fs.DesignSpec.Colors {
		if !hexColorRe.MatchString(val) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-02] color \"" + name + "\" has invalid hex value: " + val,
				Advice:  "Use a valid hex color: # followed by 3, 4, 6, or 8 hex digits",
			})
		}
	}
	return diags
}
