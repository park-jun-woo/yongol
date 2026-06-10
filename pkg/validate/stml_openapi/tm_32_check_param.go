//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-32 보조 — 단일 매핑의 소스 검사 후 대상 세그먼트를 해석 (생략형 모호·미존재 세그먼트 진단)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm32CheckParam validates one data-link-params binding: the source
// (tm32CheckParamSource) and the segment resolution. It returns the
// resolved segment name ("" when unresolvable) plus diagnostics. The
// elided form resolves only against exactly one required segment; an
// explicit SegmentName must exist in the target route (case-exact).
func tm32CheckParam(p stml.LinkParamBind, ref linkRefCtx, file, pattern string, required []string, ownRouteParams map[string]bool) (string, []diagnostic.Diagnostic) {
	diags := tm32CheckParamSource(p, ref, file, ownRouteParams)

	if p.Segment == "" && len(required) != 1 {
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-32] elided segment mapping %q is ambiguous: target page %q (resolved route %q) has %d required segments, not exactly one", p.Source, ref.Link.TargetPage, pattern, len(required)),
			Advice:  "Write the full \"<source> -> <SegmentName>\" mapping",
		})
		return "", diags
	}
	if p.Segment == "" {
		return required[0], diags
	}

	allSegments := map[string]bool{}
	for _, name := range stml.RoutePatternParams(pattern) {
		allSegments[name] = true
	}
	if !allSegments[p.Segment] {
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-32] segment %q is not in target page %q resolved route %q", p.Segment, ref.Link.TargetPage, pattern),
			Advice:  "Map to one of the target route's :Name segments (case-exact)",
		})
		return "", diags
	}
	return p.Segment, diags
}
