//ff:func feature=validate type=util control=iteration dimension=1 topic=design-structural
//ff:what checkSinglePropRefs — 단일 prop 값에서 {group.token} 참조의 resolve 여부 검사
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// checkSinglePropRefs checks all {group.token} references in a single prop value.
func checkSinglePropRefs(fs *yongol.Fullstack, compName, propName, propVal string) []diagnostic.Diagnostic {
	refs := tokenRefRe.FindAllStringSubmatch(propVal, -1)
	var diags []diagnostic.Diagnostic
	for _, ref := range refs {
		dotted := ref[1]
		if !resolveToken(fs, dotted) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-05] component \"" + compName + "\" prop \"" + propName + "\" references unresolved token: {" + dotted + "}",
				Advice:  "Ensure the referenced token exists in the DESIGN.md (colors, typography, rounded, or spacing group)",
			})
		}
	}
	return diags
}
