//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-02 — features.yaml 내 path 중복 검출

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft02DuplicatePath validates FT-02: every path in features.yaml must be unique.
func ft02DuplicatePath(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Features == nil {
		return nil
	}
	seen := make(map[string]int) // path -> first occurrence line
	var diags []diagnostic.Diagnostic
	for _, f := range fs.Features {
		if firstLine, exists := seen[f.Path]; exists {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "features.yaml",
				Line:    f.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[FT-02] duplicate path %q (first at line %d)", f.Path, firstLine),
				Advice:  "Remove the duplicate path entry from features.yaml",
			})
		} else {
			seen[f.Path] = f.Line
		}
	}
	return diags
}
