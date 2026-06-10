//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what data-redirect를 navigate 호출 라인들로 렌더링한다 (정적 경로 불변, 페이지명 참조는 응답 필드 치환 + 부재 가드)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderRedirectNavigate renders the navigate() lines for an action's
// data-redirect (page-flow Phase008). A "/"-prefixed value is a static
// path — emitted verbatim, byte-identical to the pre-Phase008 output. Any
// other value is a page-name reference: the target page's resolved route
// pattern (RedirectPattern, codegen-populated; falls back to
// "/<page-name>") gets the data-redirect-params sources substituted —
// respFields read from the 2xx response (`data`), route.<Name> from the
// useParams() variable. Optional segments are filled only when mapped;
// unmapped required segments cannot survive validation (TM-33) and are
// omitted defensively. Like the Phase003 capture guard, each substituted
// respField is guarded: a 2xx response missing it aborts the navigate and
// surfaces through the action's error state instead of emitting a path
// with "undefined" baked in. usesData reports whether the onSuccess
// handler needs the (data) parameter.
func renderRedirectNavigate(a stmlparser.ActionBlock) (lines []string, usesData bool) {
	if strings.HasPrefix(a.Redirect, "/") {
		return []string{fmt.Sprintf("navigate('%s')", a.Redirect)}, false
	}
	pattern := a.RedirectPattern
	if pattern == "" {
		pattern = "/" + a.Redirect
	}
	sources := segmentSourcesFor(pattern, a.RedirectParams)
	var parts []string
	var guarded []string // respFields needing the missing-field guard, in segment order
	seen := map[string]bool{}
	hasExpr := false
	for _, seg := range strings.Split(strings.TrimPrefix(pattern, "/"), "/") {
		if seg == "" {
			continue
		}
		if !strings.HasPrefix(seg, ":") {
			parts = append(parts, seg)
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(seg, ":"), "?")
		src, ok := sources[name]
		if !ok {
			continue
		}
		if !strings.HasPrefix(src, "route.") && !seen[src] {
			seen[src] = true
			guarded = append(guarded, src)
		}
		parts = append(parts, "${"+redirectParamExpr(src)+"}")
		hasExpr = true
	}
	errSetter := toUpperFirst(errorStateVar(a))
	for _, field := range guarded {
		lines = append(lines,
			fmt.Sprintf("if (data?.%s == null) {", field),
			fmt.Sprintf("  set%s('Unexpected response: missing %s')", errSetter, field),
			"  return",
			"}",
		)
	}
	path := "/" + strings.Join(parts, "/")
	if hasExpr {
		lines = append(lines, fmt.Sprintf("navigate(`%s`)", path))
	} else {
		lines = append(lines, fmt.Sprintf("navigate('%s')", path))
	}
	return lines, len(guarded) > 0
}
