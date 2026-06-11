//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-51 — sitemap 메뉴 렌더 항목 ≥1 인데 호스트 레이아웃 전무 (layouts/ 비고 + defaultLayout 미선언 + 모든 nav data-layout 부재) — 메뉴/브레드크럼이 렌더될 자리 없음 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm51SitemapNoLayoutHost is the inverse of TM-49 (plans/stml/sitemap
// Phase008, BUG-129): a sitemap exists and derives a menu, but no layout
// exists to host it. The sitemap's menu/breadcrumb render inside a layout's
// <slot data-outlet> (TM-44); when layouts/ is empty, manifest
// frontend.defaultLayout is unset and no nav declares data-layout, the
// pages render bare <main> with nowhere to mount the derived menu and the
// emitted Breadcrumb becomes dead code — yet validate stayed silent.
//
// The menu-rendered judgment reuses collectMenuRendered into a throwaway
// pageGraph rather than a per-node menuBlockReason walk: a data-menu="false"
// parent hides its whole subtree, so a deep child whose own menuBlockReason
// is "" still never renders. Reusing the hidden-subtree propagation keeps
// this rule's "renders a menu" judgment identical to the emitter's
// (buildSitemapMenuItems) — the Phase002 emitter/validator parity principle.
// g.Roots collects exactly the menu-rendered page nodes (recordMenuPage
// folds reason=="" data-page entries; group labels/external links skipped),
// so len(g.Roots) > 0 is the menu-renders-≥1 gate.
//
// Conditions 4/5 stay silent where TM-12 (run.go gate: defaultLayout set +
// layouts/ empty → ERROR) or TM-41 (sitemapDiags: nav data-layout +
// layouts/ empty → ERROR) already fire — TM-51 is WARNING and must not
// double-diagnose the same misconfiguration.
func tm51SitemapNoLayoutHost(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.Layouts) > 0 {
		return nil
	}
	if defaultLayoutFromManifest(fs) != "" {
		return nil
	}
	for _, nav := range fs.Sitemap.Navs {
		if nav.Layout != "" {
			return nil
		}
	}

	g := &pageGraph{
		Roots:       map[string]bool{},
		Edges:       map[string][]string{},
		InSitemap:   map[string]bool{},
		MenuBlocked: map[string]string{},
	}
	for _, nav := range fs.Sitemap.Navs {
		collectMenuRendered(nav.Items, 1, "", fs.STMLPages, g)
	}
	if len(g.Roots) == 0 {
		return nil
	}

	return []diagnostic.Diagnostic{{
		File:    fs.Sitemap.FileName,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[TM-51] %s declares %d menu-rendered entries but no layout exists to host the menu/breadcrumb — layouts/ is empty, manifest frontend.defaultLayout is unset and no sitemap nav declares data-layout, so the derived menu and breadcrumb never render", fs.Sitemap.FileName, len(g.Roots)),
		Advice:  "Create layouts/<name>.html with a <slot data-outlet> and assign it via manifest frontend.defaultLayout or <nav data-sitemap data-layout=\"<name>\">",
	}}
}
