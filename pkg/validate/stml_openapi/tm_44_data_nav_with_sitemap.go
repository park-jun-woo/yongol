//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-44 — sitemap 존재 시 레이아웃 data-nav 잔존 (단일 진실 위반, ERROR — 메뉴는 sitemap.html 로 이동)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm44DataNavWithSitemap rejects data-nav in any layout HTML while
// frontend/sitemap.html exists (plans/stml/sitemap Phase003, DESIGN §4.9):
// from Phase003 on the layout menu derives from the sitemap, so a
// surviving data-nav is a second, drifting menu truth — coexistence is an
// ERROR, not a WARNING, because tolerated drift contradicts the validate
// philosophy. Without a sitemap the data-nav path stays fully supported
// (TM-36 keeps resolving its targets); the caller gates on fs.Sitemap.
func tm44DataNavWithSitemap(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, l := range fs.Layouts {
		if len(l.NavItems) == 0 {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    l.File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-44] layout %q still declares data-nav while frontend/sitemap.html exists — the menu's single source of truth is the sitemap (메뉴는 sitemap.html 로 이동)", l.Name),
			Advice:  "Remove the data-nav entries from the layout and declare the menu structure in frontend/sitemap.html (메뉴는 sitemap.html 로 이동) — the layout keeps only the menu position",
		})
	}
	return diags
}
