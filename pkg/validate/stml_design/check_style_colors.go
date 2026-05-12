//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what checkStyleColors — style 속성 값에서 DESIGN.md 토큰과 일치하는 하드코딩 hex 색상 검출
package stml_design

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkStyleColors checks a style attribute value for hardcoded hex colors that
// match DESIGN.md token values.
func checkStyleColors(style, filename string, hexToToken map[string]string, diags *[]diagnostic.Diagnostic) {
	matches := hexColorRe.FindAllString(style, -1)
	seen := make(map[string]bool)
	for _, m := range matches {
		lower := strings.ToLower(m)
		if seen[lower] {
			continue
		}
		seen[lower] = true
		if tokenName, ok := hexToToken[lower]; ok {
			*diags = append(*diags, diagnostic.Diagnostic{
				File:    filename,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XVM-05] inline style contains hardcoded color %s which matches DESIGN.md token %q", m, tokenName),
				Advice:  fmt.Sprintf("Use Tailwind class with token %q (e.g. bg-%s, text-%s) instead of inline style", tokenName, tokenName, tokenName),
			})
		}
	}
}
