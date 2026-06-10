//ff:func feature=stml-gen type=util control=sequence
//ff:what 액션의 에러 메시지 useState 변수명을 도출한다 (data-on-error 유무 무관 상시)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// errorStateVar returns the error-message state variable name for an action
// (e.g. "Login" → "loginError"). Always non-empty since page-flow Phase004:
// every action keeps an error state so a rejected mutation is never silent,
// whether or not data-on-error is declared (BUG-113 (2) — declaring
// data-on-error only decides the display element/position).
func errorStateVar(a stmlparser.ActionBlock) string {
	return toLowerFirst(a.OperationID) + "Error"
}
