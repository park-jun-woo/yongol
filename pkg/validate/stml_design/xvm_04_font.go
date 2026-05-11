//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-04 — STML font 클래스의 이름이 DESIGN.md typography fontFamily에 없음 (WARNING)
package stml_design

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xvm04Font checks that font token names in STML classes match a typography
// fontFamily value in DESIGN.md (case-insensitive comparison).
func xvm04Font(fs *yongol.Fullstack, tokens pageTokenRefs, ovr overrideSet) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Typography) == 0 {
		return nil
	}

	// Build set of known font family names (lowercased)
	knownFonts := make(map[string]bool)
	for _, tt := range fs.DesignSpec.Typography {
		if tt.FontFamily != "" {
			knownFonts[strings.ToLower(tt.FontFamily)] = true
		}
	}
	if len(knownFonts) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	seen := make(map[string]bool)
	for _, ref := range tokens.Fonts {
		if isOverridden(ovr, ref.File, ref.Class) {
			continue
		}
		key := ref.File + ":" + ref.Name
		if seen[key] {
			continue
		}
		if !knownFonts[strings.ToLower(ref.Name)] {
			seen[key] = true
			diags = append(diags, diagnostic.Diagnostic{
				File:    ref.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XVM-04] font %q used in class is not a recognized fontFamily in DESIGN.md typography", ref.Name),
				Advice:  fmt.Sprintf("Add a typography entry with fontFamily %q in DESIGN.md, or use a standard Tailwind font-family class", ref.Name),
			})
		}
	}
	return diags
}
