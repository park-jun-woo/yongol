//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what data-capture 바인딩을 세션 store setAuth 호출 코드로 렌더링한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderCaptureCommit renders the session-store commit for the action's
// data-capture bindings: each respField is read from the mutation response
// and written to its auth sink via useAuthStore.setAuth(token, refresh).
func renderCaptureCommit(captures []stmlparser.CaptureBind) string {
	tokenExpr := "undefined"
	refreshExpr := ""
	for _, c := range captures {
		switch c.Sink {
		case "auth.token":
			tokenExpr = "data." + c.RespField
		case "auth.refresh":
			refreshExpr = "data." + c.RespField
		}
	}
	if refreshExpr == "" {
		return fmt.Sprintf("useAuthStore.getState().setAuth(%s)", tokenExpr)
	}
	return fmt.Sprintf("useAuthStore.getState().setAuth(%s, %s)", tokenExpr, refreshExpr)
}
