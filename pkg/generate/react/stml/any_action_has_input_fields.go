//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 액션 목록에서 input 태그로 렌더되는 필드가 있는지 확인한다
package stml

import (
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// anyActionHasInputFields returns true if any action has fields that render as <Input>.
func anyActionHasInputFields(actions []stmlparser.ActionBlock) bool {
	for _, a := range actions {
		if actionHasInputField(a) {
			return true
		}
	}
	return false
}
