//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-01 — features.yaml 내 op 중복 검출

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft01DuplicateOp validates FT-01: every op in features.yaml must be unique.
func ft01DuplicateOp(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Features == nil {
		return nil
	}
	seen := make(map[string]int) // op -> first occurrence line
	var diags []diagnostic.Diagnostic
	for _, f := range fs.Features {
		if firstLine, exists := seen[f.Op]; exists {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "features.yaml",
				Line:    f.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[FT-01] duplicate op %q (first at line %d)", f.Op, firstLine),
				Advice:  "Remove the duplicate op entry from features.yaml",
			})
		} else {
			seen[f.Op] = f.Line
		}
	}
	return diags
}
