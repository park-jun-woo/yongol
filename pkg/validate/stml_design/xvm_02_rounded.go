//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-02 — STML rounded 클래스의 토큰 이름이 DESIGN.md rounded에 없음 (WARNING)
package stml_design

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xvm02Rounded checks that rounded token names in STML classes exist in DESIGN.md rounded.
func xvm02Rounded(fs *yongol.Fullstack, tokens pageTokenRefs, ovr overrideSet) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Rounded) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	seen := make(map[string]bool)
	for _, ref := range tokens.Rounded {
		if isOverridden(ovr, ref.File, ref.Class) {
			continue
		}
		key := ref.File + ":" + ref.Name
		if seen[key] {
			continue
		}
		if _, ok := fs.DesignSpec.Rounded[ref.Name]; !ok {
			seen[key] = true
			diags = append(diags, diagnostic.Diagnostic{
				File:    ref.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XVM-02] rounded token %q used in class is not defined in DESIGN.md rounded", ref.Name),
				Advice:  fmt.Sprintf("Add %q to the rounded map in DESIGN.md, or use a standard Tailwind rounded value", ref.Name),
			})
		}
	}
	return diags
}
