//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what each 내부 ChildNode 슬라이스를 순회하며 행 액션 RowMutateArg 설정을 위임한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// setRowActionArgsInEachChildren recursively visits the children of a
// data-each item template and records the call-site mutate argument on
// every action that consumes item.<Field> sources.
func setRowActionArgsInEachChildren(children []stmlparser.ChildNode, itemFieldTypes map[string]string, pathParamTypes map[string]map[string]string) {
	for i := range children {
		setRowActionArgForEachChild(&children[i], itemFieldTypes, pathParamTypes)
	}
}
