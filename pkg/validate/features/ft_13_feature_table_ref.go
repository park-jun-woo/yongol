//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-structural
//ff:what FT-13 — feature의 table 값이 tables에 정의되지 않으면 ERROR

package features

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ft13FeatureTableRef validates FT-13: when a feature specifies a table field,
// that table must exist in the tables section.
func ft13FeatureTableRef(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.FeatureTables == nil {
		return nil
	}
	if fs.Features == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, f := range fs.Features {
		if f.Table == "" {
			continue
		}
		if _, ok := fs.FeatureTables[f.Table]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "features.yaml",
				Line:    f.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[FT-13] feature %q references table %q but %q is not defined in tables", f.Op, f.Table, f.Table),
				Advice:  fmt.Sprintf("Add %q to the tables section in features.yaml or correct the table field in feature %q", f.Table, f.Op),
			})
		}
	}
	return diags
}
