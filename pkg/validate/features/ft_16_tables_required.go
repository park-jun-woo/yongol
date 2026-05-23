//ff:func feature=validate type=rule control=sequence topic=features-structural
//ff:what FT-16 — features.yaml에 tables 섹션이 없으면 ERROR

package features

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft16TablesRequired validates FT-16: features.yaml must contain a non-empty
// tables section.
func ft16TablesRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.FeatureTables) > 0 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "features.yaml",
		Line:    0,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: `[FT-16] features.yaml missing required "tables" section`,
		Advice:  "Add a tables section defining your data model relationships",
	}}
}
