//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-17 — feature에 table 필드가 없으면 ERROR

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft17FeatureTableRequired validates FT-17: every feature must have a non-empty
// table field.
func ft17FeatureTableRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Features == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, f := range fs.Features {
		if f.Table == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "features.yaml",
				Line:    f.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[FT-17] feature %q missing required \"table\" field", f.Op),
				Advice:  fmt.Sprintf("Add a table field to feature %q in features.yaml", f.Op),
			})
		}
	}
	return diags
}
