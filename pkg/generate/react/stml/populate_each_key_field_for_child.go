//ff:func feature=stml-gen type=util control=selection
//ff:what 단일 ChildNode에서 fetch 기반 EachBlock KeyField 설정을 처리한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func populateEachKeyFieldForChild(ch *stmlparser.ChildNode, raif map[string]map[string]map[string]bool) {
	switch ch.Kind {
	case "fetch":
		opID := ch.Fetch.OperationID
		setEachKeyFieldsInFetch(ch.Fetch, opID, raif)
	case "static":
		if ch.Static != nil {
			populateEachKeyFieldsInChildren(ch.Static.Children, raif)
		}
	case "state":
		if ch.State != nil {
			populateEachKeyFieldsInChildren(ch.State.Children, raif)
		}
	}
}
