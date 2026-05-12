//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 페이지 Children 트리에서 fetch 기반 EachBlock KeyField를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func populateEachKeyFieldsInChildren(children []stmlparser.ChildNode, raif map[string]map[string]map[string]bool) {
	for i := range children {
		populateEachKeyFieldForChild(&children[i], raif)
	}
}
