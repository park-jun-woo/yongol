//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what data-capture 바인딩을 토큰 부재 방어 가드 + setAuth/setClaim 커밋 라인들로 렌더링한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderCaptureCommit renders the session-store commit lines for the action's
// data-capture bindings: token/refresh respFields are read from the mutation
// response and written via useAuthStore.setAuth(token, refresh); each
// auth.claims.<name> binding (plans/stml/sitemap Phase005) commits via
// setClaim('<name>', String(...)) after the token commit — claims-only
// captures (cookie mode) emit no setAuth call at all.
//
// Defensive commit (BUG-113 (3)): when an auth.token capture is declared, a
// guard precedes the commit — a 2xx response missing the token field (schema
// marks it optional, and a contract-violating server can omit it) aborts the
// commit and any redirect instead of storing undefined. The violation is
// surfaced through the action's data-on-error state when declared (errVar
// non-empty), else via console.error. auth.refresh stays optional — only the
// token gates the commit. A claim missing from the response is skipped (the
// menu simply stays role-hidden) instead of storing "undefined".
func renderCaptureCommit(captures []stmlparser.CaptureBind, errVar string) []string {
	tokenField, refreshField, claims := splitCaptureBinds(captures)
	tokenExpr := "undefined"
	var lines []string
	if tokenField != "" {
		tokenExpr = "data." + tokenField
		surface := fmt.Sprintf("console.error('Unexpected response: missing %s')", tokenField)
		if errVar != "" {
			surface = fmt.Sprintf("set%s('Unexpected response: missing %s')", toUpperFirst(errVar), tokenField)
		}
		lines = append(lines,
			fmt.Sprintf("if (data?.%s == null) {", tokenField),
			"  "+surface,
			"  return",
			"}",
		)
	}
	if tokenField != "" || refreshField != "" {
		if refreshField == "" {
			lines = append(lines, fmt.Sprintf("useAuthStore.getState().setAuth(%s)", tokenExpr))
		} else {
			lines = append(lines, fmt.Sprintf("useAuthStore.getState().setAuth(%s, data.%s)", tokenExpr, refreshField))
		}
	}
	for _, c := range claims {
		name, _ := stmlparser.ClaimsSinkName(c.Sink)
		lines = append(lines,
			fmt.Sprintf("if (data?.%s != null) {", c.RespField),
			fmt.Sprintf("  useAuthStore.getState().setClaim('%s', String(data.%s))", name, c.RespField),
			"}",
		)
	}
	return lines
}
