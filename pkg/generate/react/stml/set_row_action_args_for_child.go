//ff:func feature=stml-gen type=util control=selection
//ff:what 단일 ChildNode의 종류에 따라 행 액션 RowMutateArg 설정을 분기한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setRowActionArgsForChild(ch *stmlparser.ChildNode, opID string, itemTypes map[string]map[string]map[string]string, pathParamTypes map[string]map[string]string) {
	switch ch.Kind {
	case "each":
		setRowActionArgsInEach(ch.Each, opID, itemTypes, pathParamTypes)
	case "fetch":
		nestedOpID := ch.Fetch.OperationID
		setRowActionArgsInFetch(ch.Fetch, nestedOpID, itemTypes, pathParamTypes)
	case "static":
		if ch.Static != nil {
			setRowActionArgsInChildren(ch.Static.Children, opID, itemTypes, pathParamTypes)
		}
	case "state":
		if ch.State != nil {
			setRowActionArgsInChildren(ch.State.Children, opID, itemTypes, pathParamTypes)
		}
	}
}
