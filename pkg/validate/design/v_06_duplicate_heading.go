//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-structural
//ff:what V-06 — Markdown body에 중복 ## 섹션 헤딩 검증
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// v06DuplicateHeading validates that no ## heading is duplicated in the body.
func v06DuplicateHeading(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	seen := make(map[string]bool)
	for _, h := range fs.DesignSpec.Headings {
		if seen[h] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[V-06] duplicate ## heading: \"" + h + "\"",
				Advice:  "Remove or rename the duplicate section heading",
			})
		}
		seen[h] = true
	}
	return diags
}
