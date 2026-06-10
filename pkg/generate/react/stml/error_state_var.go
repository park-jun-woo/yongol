//ff:func feature=stml-gen type=util control=sequence
//ff:what data-on-error 액션의 에러 메시지 useState 변수명을 도출한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// errorStateVar returns the error-message state variable name for an action
// with a data-on-error slot (e.g. "Login" → "loginError"). Empty when the
// action has no data-on-error element.
func errorStateVar(a stmlparser.ActionBlock) string {
	if !a.OnErrorNode {
		return ""
	}
	return toLowerFirst(a.OperationID) + "Error"
}
