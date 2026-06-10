//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what data-capture 바인딩을 토큰 부재 방어 가드 + 세션 store setAuth 호출 라인들로 렌더링한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderCaptureCommit renders the session-store commit lines for the action's
// data-capture bindings: each respField is read from the mutation response
// and written to its auth sink via useAuthStore.setAuth(token, refresh).
//
// Defensive commit (BUG-113 (3)): when an auth.token capture is declared, a
// guard precedes the commit — a 2xx response missing the token field (schema
// marks it optional, and a contract-violating server can omit it) aborts the
// commit and any redirect instead of storing undefined. The violation is
// surfaced through the action's data-on-error state when declared (errVar
// non-empty), else via console.error. auth.refresh stays optional — only the
// token gates the commit.
func renderCaptureCommit(captures []stmlparser.CaptureBind, errVar string) []string {
	tokenField := ""
	refreshExpr := ""
	for _, c := range captures {
		switch c.Sink {
		case "auth.token":
			tokenField = c.RespField
		case "auth.refresh":
			refreshExpr = "data." + c.RespField
		}
	}
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
	if refreshExpr == "" {
		return append(lines, fmt.Sprintf("useAuthStore.getState().setAuth(%s)", tokenExpr))
	}
	return append(lines, fmt.Sprintf("useAuthStore.getState().setAuth(%s, %s)", tokenExpr, refreshExpr))
}
