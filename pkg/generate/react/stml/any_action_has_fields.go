//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 액션 목록에서 폼 필드를 가진 액션이 있는지 확인한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// anyActionHasFields returns true if any action has form fields.
func anyActionHasFields(actions []stmlparser.ActionBlock) bool {
	for _, a := range actions {
		if len(a.Fields) > 0 {
			return true
		}
	}
	return false
}
