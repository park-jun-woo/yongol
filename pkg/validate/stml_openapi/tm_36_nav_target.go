//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-36 — 레이아웃 data-nav 대상 해석 불가 검출 (모든 레이아웃 × NavItem 순회, ERROR)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm36NavTarget checks every layout data-nav target (page-flow Phase010,
// BUG-114 (4)) under the same dual rule as data-redirect (Phase008): a
// "/"-prefixed value is a static path matched against every page's
// resolved route patterns ("/" is always allowed — the index route is
// emitted), any other value is a page-name reference (STML filename
// without .html). A page-name reference additionally must resolve to a
// route without a required parameter segment — a static menu link has no
// context to fill it (parameterized navigation is data-link territory,
// Phase007). A broken menu entry would render on every page using the
// layout, so each violation is an ERROR.
func tm36NavTarget(layouts []stml.LayoutSpec, pages []stml.PageSpec) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, l := range layouts {
		for _, item := range l.NavItems {
			diags = append(diags, tm36NavItemDiags(l, item, pages)...)
		}
	}
	return diags
}
