//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-32 보조 — 단일 동적 그룹의 link-params 를 구문·소스·세그먼트·필수 충족 순으로 검사 (tm32CheckLink 의 사이트맵 판)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm32CheckSitemapGroup validates one dynamic menu group's
// data-link-params against its target page's resolved route — the
// tm32CheckLink sequence with the group's each item schema as the item
// context (every emitted item is a data-each row, so item.* is always in
// scope) and the route.* rejection of tm32SitemapParam in place of the
// own-route check. A syntax error in the raw attribute is reported alone —
// the bindings are unusable until it is fixed.
func tm32CheckSitemapGroup(e sitemapEntry, file string, patterns map[string]string, raif map[string]map[string]map[string]bool) []diagnostic.Diagnostic {
	pattern, ok := patterns[e.Node.Link]
	if !ok {
		return nil // TM-31/TM-48 own the missing-target findings
	}

	var params []stml.LinkParamBind
	if e.Node.LinkParamsRaw != "" {
		var err error
		params, err = stml.ParseLinkParams(e.Node.LinkParamsRaw)
		if err != nil {
			return []diagnostic.Diagnostic{{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-32] dynamic menu group data-link-params %q at %s: %v", e.Node.LinkParamsRaw, e.Path, err),
				Advice:  "Use \"item.<Field> -> <SegmentName>\" pairs (comma-separated); the \"-> <SegmentName>\" part may be elided only against a single required segment",
			}}
		}
	}

	var required []string
	for _, seg := range strings.Split(pattern, "/") {
		if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
			required = append(required, strings.TrimPrefix(seg, ":"))
		}
	}

	ref := linkRefCtx{
		Link:       &stml.LinkRef{TargetPage: e.Node.Link, ParamsRaw: e.Node.LinkParamsRaw},
		ItemFields: raif[e.Node.Fetch][e.Node.Each],
		InEach:     true, // the emitted items are the data-each rows — item.* is always in scope
	}
	var diags []diagnostic.Diagnostic
	mapped := map[string]bool{}
	for _, p := range params {
		seg, pd := tm32SitemapParam(p, ref, e.Path, file, pattern, required)
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
			Message: fmt.Sprintf("[TM-32] required segment :%s of target page %q (resolved route %q) is not mapped by the dynamic menu group at %s — the emitted item links would miss a path segment", name, e.Node.Link, pattern, e.Path),
			Advice:  fmt.Sprintf("Add \"item.<Field> -> %s\" to the group's data-link-params", name),
		})
	}
	return diags
}
