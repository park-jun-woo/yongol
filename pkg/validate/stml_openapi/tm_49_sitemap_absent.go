//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-49 — frontend ON + 페이지 ≥1 인데 sitemap.html 부재 — 사이트 구조 미선언 안내 (WARNING)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm49SitemapAbsent surfaces the undeclared-site-structure state
// (plans/stml/sitemap Phase001, DESIGN §4.1): the frontend is ON with at
// least one STML page but frontend/sitemap.html does not exist, so menu /
// breadcrumb derivation and reachability (orphan page) validation are
// inactive. A central file's absence is itself detectable — this is the
// rule that keeps a sitemap-less gozhip-style spec from passing without a
// word (one axis of closing BUG-122; the other is Phase002's TM-43).
func tm49SitemapAbsent(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Sitemap != nil {
		return nil
	}
	if !frontendEnabled(fs) || len(fs.STMLPages) == 0 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "frontend/sitemap.html",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[TM-49] no frontend/sitemap.html — the site structure is undeclared, so menu/breadcrumb derivation and reachability (orphan page) validation are inactive",
		Advice:  "Declare the site tree in frontend/sitemap.html: <nav data-sitemap> blocks with nested <ul>/<li data-page> entries (one li per page, document order = menu order)",
	}}
}
