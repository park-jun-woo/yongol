//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-32 — data-link-params가 대상 라우트의 필수 세그먼트를 충족하지 못함 (ERROR, 구문·세그먼트·소스 검사)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm32LinkParamsUnsatisfied checks every data-link's param bindings on the
// page (page-flow Phase007): syntax (re-parse of the raw attribute, the
// ParseCapture/TM-20 split), segment-name existence in the target page's
// resolved route, required-segment coverage, item.<Field> source legality
// (TM-30 infrastructure: inside data-each + item schema), route.<Name>
// source existence in this page's own resolved route (TM-27
// infrastructure), and the elided form's single-required-segment rule.
// Unmapped optional segments are legal (silent). A target page missing
// from the set stays silent — TM-31 owns that diagnostic.
func tm32LinkParamsUnsatisfied(page stml.PageSpec, pages []stml.PageSpec, raif map[string]map[string]map[string]bool) []diagnostic.Diagnostic {
	var refs []linkRefCtx
	collectLinkRefs(page.Children, "", nil, false, raif, &refs)
	if len(refs) == 0 {
		return nil
	}
	patterns := map[string]string{}
	for _, p := range pages {
		if rp := stml.RoutePaths(p); len(rp) > 0 {
			patterns[p.Name] = rp[0]
		}
	}
	ownRouteParams := map[string]bool{}
	for _, pattern := range stml.RoutePaths(page) {
		for _, name := range stml.RoutePatternParams(pattern) {
			ownRouteParams[name] = true
		}
	}
	var diags []diagnostic.Diagnostic
	for _, ref := range refs {
		diags = append(diags, tm32CheckLink(ref, page.FileName, patterns, ownRouteParams)...)
	}
	return diags
}
