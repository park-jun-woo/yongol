//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-10 — has_many가 참조하는 테이블이 tables에 정의되지 않으면 ERROR

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft10HasManyRef validates FT-10: every has_many reference must point to a
// table defined in the tables section.
func ft10HasManyRef(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for parent, td := range fs.FeatureTables {
		for _, child := range td.HasMany {
			if _, ok := fs.FeatureTables[child]; !ok {
				diags = append(diags, diagnostic.Diagnostic{
					File:    "features.yaml",
					Line:    0,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[FT-10] table %q has_many references %q but %q is not defined in tables", parent, child, child),
					Advice:  fmt.Sprintf("Add %q to the tables section in features.yaml or remove it from %q has_many", child, parent),
				})
			}
		}
	}
	return diags
}
