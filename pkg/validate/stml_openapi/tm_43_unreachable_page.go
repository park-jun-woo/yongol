//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-43 — 루트에서 간선으로 도달 불가능한 고아 페이지 (원인 분류 포함, WARNING, sitemap 존재 시에만)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm43UnreachablePage warns on every STML page the user cannot actually
// reach (plans/stml/sitemap Phase002, DESIGN §4.10 — the rule closing
// BUG-122's gozhip scenario together with TM-49): BFS from the roots
// (index/entry pages plus menu-rendered sitemap entries) over the
// data-link/data-redirect edges plus the Phase004 breadcrumb up-edges
// (a reachable page's breadcrumb links its MenuRenderable sitemap
// ancestors — DESIGN §4.10 edge (d)). Listing in the sitemap is a node, not an
// edge — a listed entry that does not render in the menu (required route
// param, depth > 2, data-menu="false") still needs an incoming link, so
// listing alone never silences the warning. Active only when a sitemap
// exists: it is the place where roots and data-entry opt-outs are
// declared, so a sitemap-less reachability check could not avoid flagging
// public entry pages (TM-49 covers the absence instead). The message
// classifies the cause (listed-but-not-menu-rendered vs not listed,
// no incoming edge vs only unreachable sources) to point at the fix.
func tm43UnreachablePage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Sitemap == nil || len(fs.STMLPages) == 0 {
		return nil
	}
	g := buildPageGraph(fs)
	reached := reachablePages(g)
	hasIncoming := map[string]bool{}
	for _, targets := range g.Edges {
		for _, t := range targets {
			hasIncoming[t] = true
		}
	}

	var diags []diagnostic.Diagnostic
	for _, p := range fs.STMLPages {
		if reached[p.Name] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    p.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[TM-43] page %q is unreachable: %s", p.Name, tm43Cause(g, p.Name, hasIncoming[p.Name])),
			Advice:  "Give it an incoming edge: declare data-link on a reachable page (e.g. a list row), data-redirect on an action, make its sitemap entry menu-renderable (a reachable descendant's breadcrumb then links up to it), or put it in a <nav data-sitemap data-entry> block if it is a public entry point",
		})
	}
	return diags
}
