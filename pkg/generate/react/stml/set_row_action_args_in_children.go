//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리에서 EachBlock의 행 액션에 RowMutateArg를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setRowActionArgsInChildren(children []stmlparser.ChildNode, opID string, itemTypes map[string]map[string]map[string]string, pathParamTypes map[string]map[string]string) {
	for i := range children {
		setRowActionArgsForChild(&children[i], opID, itemTypes, pathParamTypes)
	}
}
