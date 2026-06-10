//ff:func feature=stml-gen type=util control=selection
//ff:what each 내부 단일 ChildNode의 종류에 따라 행 액션 RowMutateArg를 기록한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// setRowActionArgForEachChild handles one ChildNode inside a data-each item
// template. Actions without item params keep the hoisted-closure path
// (RowMutateArg stays empty).
func setRowActionArgForEachChild(ch *stmlparser.ChildNode, itemFieldTypes map[string]string, pathParamTypes map[string]map[string]string) {
	switch ch.Kind {
	case "action":
		if actionHasItemParam(*ch.Action) {
			ch.Action.RowMutateArg = renderRowMutateArg(*ch.Action, pathParamTypes, itemFieldTypes)
		}
	case "static":
		if ch.Static != nil {
			setRowActionArgsInEachChildren(ch.Static.Children, itemFieldTypes, pathParamTypes)
		}
	case "state":
		if ch.State != nil {
			setRowActionArgsInEachChildren(ch.State.Children, itemFieldTypes, pathParamTypes)
		}
	}
}
