//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 액션 목록에 Login operationID가 포함되어 있는지 확인한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// hasLoginAction returns true if any action in the slice has operationID "Login".
func hasLoginAction(actions []stmlparser.ActionBlock) bool {
	for _, a := range actions {
		if isLoginAction(a.OperationID) {
			return true
		}
	}
	return false
}
