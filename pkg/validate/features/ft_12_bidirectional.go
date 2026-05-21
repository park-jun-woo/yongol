//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-12 — has_many 관계에 대응하는 belongs_to가 없으면 WARN

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft12Bidirectional validates FT-12: when table A has_many B, table B should
// belongs_to A. Missing reverse direction emits a warning.
func ft12Bidirectional(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for parent, td := range fs.FeatureTables {
		for _, child := range td.HasMany {
			childDef, ok := fs.FeatureTables[child]
			if !ok {
				// FT-10 already catches this; skip here.
				continue
			}
			if !containsStr(childDef.BelongsTo, parent) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    "features.yaml",
					Line:    0,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: fmt.Sprintf("[FT-12] table %q has_many %q but %q does not belongs_to %q", parent, child, child, parent),
					Advice:  fmt.Sprintf("Add %q to %q belongs_to in features.yaml", parent, child),
				})
			}
		}
	}
	return diags
}

// containsStr returns true if slice contains the target string.
func containsStr(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
