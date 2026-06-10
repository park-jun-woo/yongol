//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-27 — 페이지가 소비하는 route.<Name>이 해석된 라우트 경로에 동명 세그먼트로 없음 (ERROR, 런타임 항상 undefined)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm27RouteParamMissing checks that every route.<Name> the page consumes
// appears as a same-named segment (":Name" or ":Name?") in the page's
// resolved route patterns (stml.RoutePaths). A missing segment means
// react-router never fills the useParams() key, so the param is always
// undefined at runtime (BUG-112's NaN calls). Derived routes are built
// from the consumed set and satisfy the invariant by construction — the
// rule still runs on them as a regression guard against a second ":id"
// hardcoding in the emitter; the practical firing surface is explicit
// data-route pages. Names are compared case-exactly: useParams() keys are
// case-sensitive, so a loose match would be a false pass.
func tm27RouteParamMissing(p stml.PageSpec) []diagnostic.Diagnostic {
	consumed := stml.ConsumedRouteParams(p)
	if len(consumed) == 0 {
		return nil
	}
	patterns := stml.RoutePaths(p)
	declared := map[string]bool{}
	for _, pattern := range patterns {
		for _, name := range stml.RoutePatternParams(pattern) {
			declared[name] = true
		}
	}
	var diags []diagnostic.Diagnostic
	for _, name := range consumed {
		if declared[name] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    p.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-27] page consumes route.%s but its resolved route %q has no :%s segment — the param is always undefined at runtime", name, strings.Join(patterns, ", "), name),
			Advice:  fmt.Sprintf("Add `:%s` to data-route or rename the param source", name),
		})
	}
	return diags
}
