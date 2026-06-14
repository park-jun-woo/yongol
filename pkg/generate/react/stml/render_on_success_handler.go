//ff:func feature=stml-gen type=generator control=sequence
//ff:what 선언 기반으로 mutation onSuccess 핸들러를 렌더링한다 (캡처 커밋 또는 removeQueries/invalidate 후 navigate 결합)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderOnSuccessHandler renders the mutation onSuccess handler. The body is
// composed from the STML flow declarations:
//  1. data-capture (bearer mode) → commit response fields to the session
//     store, guarded against a 2xx response missing the token field — the
//     guard returns early so a later redirect is also aborted (BUG-113 (3)).
//     A capture action drives its own navigation, so it never touches the
//     query cache.
//  2. otherwise (a plain mutation) → refresh the affected queries
//     (renderInvalidateExpr) — a delete instead drops its own resource GET
//     with removeQueries so a navigate away never refetches a 404 (BUG-132
//     132-2) — and then navigate to the data-redirect target. invalidate
//     and navigate are *combined*, not exclusive: TM-57 makes the redirect
//     required, so the list refreshes and the screen moves together (BUG-132
//     132-1).
//
// The error-state reset lives in onMutate (page-flow Phase004) — every
// (re)submission clears the previous message, so onSuccess no longer resets.
func renderOnSuccessHandler(a stmlparser.ActionBlock, captures []stmlparser.CaptureBind, invalidateOps, removeOps []string) string {
	var lines []string
	param := "()"

	if len(captures) > 0 {
		param = "(data)"
		lines = append(lines, renderCaptureCommit(captures, errorStateVar(a))...)
	} else {
		if rem := renderRemoveQueriesExpr(removeOps); rem != "" {
			lines = append(lines, rem)
		}
		if inv := renderInvalidateExpr(invalidateOps); inv != "" {
			lines = append(lines, inv)
		}
	}
	if a.Redirect != "" {
		nav, usesData := renderRedirectNavigate(a)
		if usesData {
			param = "(data)"
		}
		lines = append(lines, nav...)
	}
	return fmt.Sprintf(`    onSuccess: %s => {
      %s
    },
`, param, strings.Join(lines, "\n      "))
}
