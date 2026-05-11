//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-03 — STML spacing 클래스의 토큰 이름이 DESIGN.md spacing에 없음 (WARNING)
package stml_design

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xvm03Spacing checks that spacing token names in STML classes exist in DESIGN.md spacing.
func xvm03Spacing(fs *yongol.Fullstack, tokens pageTokenRefs, ovr overrideSet) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Spacing) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	seen := make(map[string]bool)
	for _, ref := range tokens.Spacing {
		if isOverridden(ovr, ref.File, ref.Class) {
			continue
		}
		key := ref.File + ":" + ref.Name
		if seen[key] {
			continue
		}
		if _, ok := fs.DesignSpec.Spacing[ref.Name]; !ok {
			seen[key] = true
			diags = append(diags, diagnostic.Diagnostic{
				File:    ref.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XVM-03] spacing token %q used in class is not defined in DESIGN.md spacing", ref.Name),
				Advice:  fmt.Sprintf("Add %q to the spacing map in DESIGN.md, or use a standard Tailwind spacing value", ref.Name),
			})
		}
	}
	return diags
}
