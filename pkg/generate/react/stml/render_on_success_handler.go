//ff:func feature=stml-gen type=generator control=sequence
//ff:what 선언 기반 3분기(캡처 커밋/리다이렉트/invalidate)로 mutation onSuccess 핸들러를 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderOnSuccessHandler renders the mutation onSuccess handler. The body is
// decided by the STML flow declarations:
//  1. data-capture (bearer mode) → commit response fields to the session
//     store, guarded against a 2xx response missing the token field — the
//     guard returns early so a later redirect is also aborted (BUG-113 (3))
//  2. data-redirect → navigate (combinable with 1): a static "/"-prefixed
//     path verbatim, or a page-name reference with data-redirect-params
//     response fields substituted into the target route, each guarded
//     against a missing field (renderRedirectNavigate, page-flow Phase008)
//  3. neither → the default invalidateQueries()/data-invalidates path
//
// The error-state reset lives in onMutate (page-flow Phase004) — every
// (re)submission clears the previous message, so onSuccess no longer resets.
func renderOnSuccessHandler(a stmlparser.ActionBlock, captures []stmlparser.CaptureBind, fetchOps []string) string {
	var lines []string
	param := "()"
	if len(captures) > 0 {
		param = "(data)"
		lines = append(lines, renderCaptureCommit(captures, errorStateVar(a))...)
	}
	if a.Redirect != "" {
		nav, usesData := renderRedirectNavigate(a)
		if usesData {
			param = "(data)"
		}
		lines = append(lines, nav...)
	}
	if len(captures) == 0 && a.Redirect == "" {
		lines = append(lines, renderInvalidateExpr(fetchOps))
	}
	return fmt.Sprintf(`    onSuccess: %s => {
      %s
    },
`, param, strings.Join(lines, "\n      "))
}
