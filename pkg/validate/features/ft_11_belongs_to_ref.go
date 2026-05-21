//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-11 — belongs_to가 참조하는 테이블이 tables에 정의되지 않으면 ERROR

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft11BelongsToRef validates FT-11: every belongs_to reference must point to a
// table defined in the tables section.
func ft11BelongsToRef(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.FeatureTables == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for child, td := range fs.FeatureTables {
		for _, parent := range td.BelongsTo {
			if _, ok := fs.FeatureTables[parent]; !ok {
				diags = append(diags, diagnostic.Diagnostic{
					File:    "features.yaml",
					Line:    0,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[FT-11] table %q belongs_to references %q but %q is not defined in tables", child, parent, parent),
					Advice:  fmt.Sprintf("Add %q to the tables section in features.yaml or remove it from %q belongs_to", parent, child),
				})
			}
		}
	}
	return diags
}
