//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 단일 액션의 필드 중 Input으로 렌더되는 것이 있는지 확인한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// actionHasInputField returns true if the action has at least one field
// that renders as <Input> (i.e. not a data-component reference).
func actionHasInputField(a stmlparser.ActionBlock) bool {
	for _, f := range a.Fields {
		if !strings.HasPrefix(f.Tag, "data-component:") {
			return true
		}
	}
	return false
}
