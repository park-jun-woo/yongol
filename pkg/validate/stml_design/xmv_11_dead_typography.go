//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XMV-11 — DESIGN.md typography 토큰이 STML에서 미참조 (WARNING)
package stml_design

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmv11DeadTypography detects DESIGN.md typography tokens whose fontFamily is
// not referenced by any STML font-* class.
func xmv11DeadTypography(fs *yongol.Fullstack, tokens pageTokenRefs) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Typography) == 0 {
		return nil
	}

	// Build set of referenced font family names (lowercased)
	usedFonts := make(map[string]bool)
	for _, ref := range tokens.Fonts {
		usedFonts[strings.ToLower(ref.Name)] = true
	}

	var diags []diagnostic.Diagnostic
	names := sortedTypoKeys(fs.DesignSpec.Typography)
	for _, name := range names {
		tt := fs.DesignSpec.Typography[name]
		if tt.FontFamily == "" {
			continue
		}
		if !usedFonts[strings.ToLower(tt.FontFamily)] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XMV-11] typography token %q (fontFamily %q) is defined in DESIGN.md but not referenced in any STML page", name, tt.FontFamily),
				Advice:  fmt.Sprintf("Use font-%s in an STML class or remove the typography entry if unused", strings.ToLower(tt.FontFamily)),
			})
		}
	}
	return diags
}

// sortedTypoKeys returns the keys of a TypographyToken map in sorted order.
func sortedTypoKeys(m map[string]design.TypographyToken) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
