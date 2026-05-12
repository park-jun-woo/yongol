//ff:func feature=generate type=util control=selection
//ff:what 단일 ChildNode에서 field-less 액션을 추출하거나 자식 노드를 재귀 탐색한다

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// appendFieldlessFromChild handles a single ChildNode: if it is a field-less
// action, add its operationId; otherwise recurse into its children.
func appendFieldlessFromChild(ch stmlparser.ChildNode, result map[string]bool) {
	switch ch.Kind {
	case "action":
		if ch.Action != nil && len(ch.Action.Fields) == 0 {
			result[ch.Action.OperationID] = true
		}
	case "fetch":
		if ch.Fetch != nil {
			collectNestedFieldlessActions(ch.Fetch.Children, result)
		}
	case "state":
		if ch.State != nil {
			collectNestedFieldlessActions(ch.State.Children, result)
		}
	case "static":
		if ch.Static != nil {
			collectNestedFieldlessActions(ch.Static.Children, result)
		}
	case "each":
		if ch.Each != nil {
			collectNestedFieldlessActions(ch.Each.Children, result)
		}
	}
}
