//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-01 — STML 색상 클래스의 토큰 이름이 DESIGN.md colors에 없음 (WARNING)
package stml_design

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xvm01Color checks that color token names in STML classes exist in DESIGN.md colors.
func xvm01Color(fs *yongol.Fullstack, tokens pageTokenRefs, ovr overrideSet) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Colors) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	seen := make(map[string]bool) // deduplicate per file+name
	for _, ref := range tokens.Colors {
		if isOverridden(ovr, ref.File, ref.Class) {
			continue
		}
		key := ref.File + ":" + ref.Name
		if seen[key] {
			continue
		}
		if _, ok := fs.DesignSpec.Colors[ref.Name]; !ok {
			seen[key] = true
			diags = append(diags, diagnostic.Diagnostic{
				File:    ref.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XVM-01] color token %q used in class is not defined in DESIGN.md colors", ref.Name),
				Advice:  fmt.Sprintf("Add %q to the colors map in DESIGN.md, or use a standard Tailwind color name", ref.Name),
			})
		}
	}
	return diags
}
