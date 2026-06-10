//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-33 — data-redirect-params가 대상 라우트의 필수 세그먼트를 충족하지 못하거나 정적 경로와 모순 (ERROR)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm33RedirectParams checks an action's data-redirect-params against its
// data-redirect target (page-flow Phase008, the TM-32 twin for redirects):
// a static "/"-prefixed redirect must not declare params (segment
// substitution applies only to page-name references — contradiction),
// while a page-name redirect is checked for syntax (re-parse of the raw
// attribute, the ParseCapture/TM-20 split), respField existence in the
// operation's 2xx response schema (TM-20 infrastructure; route.<Name>
// sources are exempt), segment-name existence in the target page's
// resolved route, required-segment coverage (also with no params attr at
// all), and the elided form's single-required-segment rule. Unmapped
// optional segments are legal (silent). A target page missing from the
// set stays silent — TM-26 owns that diagnostic.
func tm33RedirectParams(a stml.ActionBlock, file string, opMap map[string]operationEntry, pages []stml.PageSpec) []diagnostic.Diagnostic {
	if a.Redirect == "" {
		return nil
	}
	if strings.HasPrefix(a.Redirect, "/") {
		if a.RedirectParamsRaw == "" {
			return nil
		}
		return []diagnostic.Diagnostic{{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-33] data-redirect-params %q on action %q contradicts the static data-redirect %q — segment substitution applies only to page-name references", a.RedirectParamsRaw, a.OperationID, a.Redirect),
			Advice:      "Replace data-redirect with a page-name reference (STML filename without .html), or drop data-redirect-params",
			OperationID: a.OperationID,
		}}
	}

	pattern := ""
	found := false
	for _, p := range pages {
		if p.Name != a.Redirect {
			continue
		}
		found = true
		if rp := stml.RoutePaths(p); len(rp) > 0 {
			pattern = rp[0]
		}
		break
	}
	if !found || pattern == "" {
		return nil // TM-26 owns the missing-target diagnostic
	}

	var params []stml.LinkParamBind
	if a.RedirectParamsRaw != "" {
		var err error
		params, err = stml.ParseRedirectParams(a.RedirectParamsRaw)
		if err != nil {
			return []diagnostic.Diagnostic{{
				File:        file,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-33] data-redirect-params %q on action %q: %v", a.RedirectParamsRaw, a.OperationID, err),
				Advice:      "Use \"<respField> -> <SegmentName>\" pairs (comma-separated; route.<Name> sources forward a current-page param); the \"-> <SegmentName>\" part may be elided only against a single required segment",
				OperationID: a.OperationID,
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
		seg, pd := tm33CheckParam(p, a, file, pattern, required, opMap)
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
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-33] required segment :%s of redirect target page %q (resolved route %q) is not mapped — the emitted navigate() would miss a path segment", name, a.Redirect, pattern),
			Advice:      fmt.Sprintf("Add \"<respField> -> %s\" to data-redirect-params on action %q", name, a.OperationID),
			OperationID: a.OperationID,
		})
	}
	return diags
}
