//ff:func feature=stml-gen type=util control=selection
//ff:what 단일 ChildNode의 종류에 따라 EachBlock KeyField를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setEachKeyFieldForChild(ch *stmlparser.ChildNode, opID string, raif map[string]map[string]map[string]bool) {
	switch ch.Kind {
	case "each":
		setKeyFieldIfHasID(ch.Each, opID, raif)
	case "fetch":
		nestedOpID := ch.Fetch.OperationID
		setEachKeyFieldsInFetch(ch.Fetch, nestedOpID, raif)
	case "static":
		if ch.Static != nil {
			setEachKeyFieldsInChildren(ch.Static.Children, opID, raif)
		}
	case "state":
		if ch.State != nil {
			setEachKeyFieldsInChildren(ch.State.Children, opID, raif)
		}
	}
}
