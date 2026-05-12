//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리에서 EachBlock의 KeyField를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setEachKeyFieldsInChildren(children []stmlparser.ChildNode, opID string, raif map[string]map[string]map[string]bool) {
	for i := range children {
		setEachKeyFieldForChild(&children[i], opID, raif)
	}
}
