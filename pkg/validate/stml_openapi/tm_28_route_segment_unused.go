//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-28 — 라우트 경로의 :Name 세그먼트를 어떤 data-param-*도 소비하지 않음 (WARNING, 죽은 세그먼트)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm28RouteSegmentUnused warns when a ":Name"/":Name?" segment of the
// page's resolved route patterns (stml.RoutePaths) is consumed by no
// data-param-* binding on the page — a dead segment, signalling drift
// between the URL design and the page implementation (BUG-112). Derived
// routes only ever contain consumed params, so the firing surface is
// explicit data-route pages; like TM-27 the comparison is case-exact
// (useParams() keys are case-sensitive).
func tm28RouteSegmentUnused(p stml.PageSpec) []diagnostic.Diagnostic {
	consumed := map[string]bool{}
	for _, name := range stml.ConsumedRouteParams(p) {
		consumed[name] = true
	}
	var diags []diagnostic.Diagnostic
	for _, pattern := range stml.RoutePaths(p) {
		for _, name := range stml.RoutePatternParams(pattern) {
			if consumed[name] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    p.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[TM-28] route %q declares segment :%s but no data-param-* on the page consumes route.%s — dead segment", pattern, name, name),
				Advice:  fmt.Sprintf("Remove the segment or bind it with `data-param-*=\"route.%s\"`", name),
			})
		}
	}
	return diags
}
