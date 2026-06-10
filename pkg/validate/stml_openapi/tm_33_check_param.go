//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-33 보조 — 단일 매핑의 respField 소스(2xx 응답 스키마, TM-20 인프라)와 대상 세그먼트를 검사한다

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm33CheckParam validates one data-redirect-params binding: the source
// (tm33CheckParamSource) and the segment resolution, returning the
// resolved segment name ("" when unresolvable) plus diagnostics. The
// elided form resolves only against exactly one required segment; an
// explicit SegmentName must exist in the target route (case-exact) — the
// TM-32 rules under the Phase007 grammar.
func tm33CheckParam(p stml.LinkParamBind, a stml.ActionBlock, file, pattern string, required []string, opMap map[string]operationEntry) (string, []diagnostic.Diagnostic) {
	diags := tm33CheckParamSource(p, a, file, opMap)

	if p.Segment == "" && len(required) != 1 {
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-33] elided segment mapping %q is ambiguous: redirect target page %q (resolved route %q) has %d required segments, not exactly one", p.Source, a.Redirect, pattern, len(required)),
			Advice:      "Write the full \"<source> -> <SegmentName>\" mapping",
			OperationID: a.OperationID,
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
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-33] segment %q is not in redirect target page %q resolved route %q", p.Segment, a.Redirect, pattern),
			Advice:      "Map to one of the target route's :Name segments (case-exact)",
			OperationID: a.OperationID,
		})
		return "", diags
	}
	return p.Segment, diags
}
