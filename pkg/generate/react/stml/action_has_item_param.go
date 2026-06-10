//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ActionBlock의 파라미터에 item.* 행 컨텍스트 소스가 있는지 확인한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// actionHasItemParam reports whether any data-param-* source of the action
// references the current data-each row (item.<Field>). Such actions receive
// their arguments at the mutate() call site instead of the hoisted
// mutationFn closure, because `item` is only in scope inside the map
// callback (page-flow Phase006).
func actionHasItemParam(a stmlparser.ActionBlock) bool {
	for _, p := range a.Params {
		if strings.HasPrefix(p.Source, "item.") {
			return true
		}
	}
	return false
}
