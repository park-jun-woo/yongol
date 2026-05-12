//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 모든 액션이 Login인지 확인한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// allLoginActions returns true if every action in the slice is a Login action.
func allLoginActions(actions []stmlparser.ActionBlock) bool {
	for _, a := range actions {
		if !isLoginAction(a.OperationID) {
			return false
		}
	}
	return true
}
