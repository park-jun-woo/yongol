//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-structural
//ff:what V-03 — typography 토큰에 fontFamily, fontSize, fontWeight 필수 검증
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// v03TypographyRequired validates that each typography token has fontFamily, fontSize, fontWeight.
func v03TypographyRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for name, tok := range fs.DesignSpec.Typography {
		if tok.FontFamily == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-03] typography \"" + name + "\" missing required field: fontFamily",
				Advice:  "Add fontFamily to the typography token",
			})
		}
		if tok.FontSize == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-03] typography \"" + name + "\" missing required field: fontSize",
				Advice:  "Add fontSize to the typography token",
			})
		}
		if tok.FontWeight == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-03] typography \"" + name + "\" missing required field: fontWeight",
				Advice:  "Add fontWeight to the typography token",
			})
		}
	}
	return diags
}
