//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-32 보조 — 동적 그룹 매핑 1건 검사: route.* 소스는 즉시 거부(메뉴엔 라우트 문맥 없음), item.* 는 tm32CheckParam 위임

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm32SitemapParam validates one dynamic-group binding. A route.<Name>
// source is rejected outright — the layout menu renders on every route,
// so no own-route segment exists to read (pages resolve route.* against
// their own pattern instead); its explicit segment still counts as mapped
// so the finding stays singular. item.* sources delegate to tm32CheckParam
// (the page judgment) with the group's item schema in scope.
func tm32SitemapParam(p stml.LinkParamBind, ref linkRefCtx, path, file, pattern string, required []string) (string, []diagnostic.Diagnostic) {
	if !strings.HasPrefix(p.Source, "route.") {
		return tm32CheckParam(p, ref, file, pattern, required, nil)
	}
	return p.Segment, []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-32] dynamic menu group param source %q at %s reads a route segment — the layout menu renders on every route, so no own-route value exists", p.Source, path),
		Advice:  "Use an item.<Field> source; route.* belongs to page-level data-link",
	}}
}
