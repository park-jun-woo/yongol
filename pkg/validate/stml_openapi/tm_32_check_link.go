//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-32 보조 — 단일 data-link의 파라미터 매핑을 구문·소스·세그먼트·필수 충족 순으로 검사

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm32CheckLink validates one collected data-link against its target
// page's resolved route pattern. Diagnostics include the resolved pattern
// so the fix is mechanical. A syntax error in the raw attribute is
// reported alone — the bindings are unusable until it is fixed.
func tm32CheckLink(ref linkRefCtx, file string, patterns map[string]string, ownRouteParams map[string]bool) []diagnostic.Diagnostic {
	lr := ref.Link
	pattern, ok := patterns[lr.TargetPage]
	if !ok {
		return nil // TM-31 owns the missing-target diagnostic
	}

	var params []stml.LinkParamBind
	if lr.ParamsRaw != "" {
		var err error
		params, err = stml.ParseLinkParams(lr.ParamsRaw)
		if err != nil {
			return []diagnostic.Diagnostic{{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-32] data-link-params %q: %v", lr.ParamsRaw, err),
				Advice:  "Use \"<source> -> <SegmentName>\" pairs (comma-separated); the \"-> <SegmentName>\" part may be elided only against a single required segment",
			}}
		}
	}

	var required []string
	for _, seg := range strings.Split(pattern, "/") {
		if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
			required = append(required, strings.TrimPrefix(seg, ":"))
		}
	}

	var diags []diagnostic.Diagnostic
	mapped := map[string]bool{}
	for _, p := range params {
		seg, pd := tm32CheckParam(p, ref, file, pattern, required, ownRouteParams)
		diags = append(diags, pd...)
		if seg != "" {
			mapped[seg] = true
		}
	}

	for _, name := range required {
		if mapped[name] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-32] required segment :%s of target page %q (resolved route %q) is not mapped — the emitted link would miss a path segment", name, lr.TargetPage, pattern),
			Advice:  fmt.Sprintf("Add \"<source> -> %s\" to data-link-params", name),
		})
	}
	return diags
}
