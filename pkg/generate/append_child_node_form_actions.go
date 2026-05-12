//ff:func feature=generate type=util control=selection
//ff:what 단일 ChildNode에서 폼 필드 액션을 추출하거나 자식 노드를 재귀 탐색한다
package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// appendChildNodeFormActions handles a single ChildNode: if it is an action
// with fields, add it; otherwise recurse into its children.
func appendChildNodeFormActions(result []actionEntry, ch stmlparser.ChildNode, seen map[string]bool) []actionEntry {
	switch ch.Kind {
	case "action":
		if ch.Action != nil && len(ch.Action.Fields) > 0 && !seen[ch.Action.OperationID] {
			seen[ch.Action.OperationID] = true
			result = append(result, toActionEntry(*ch.Action))
		}
	case "fetch":
		if ch.Fetch != nil {
			result = append(result, collectNestedFormActions(ch.Fetch.Children, seen)...)
		}
	case "state":
		if ch.State != nil {
			result = append(result, collectNestedFormActions(ch.State.Children, seen)...)
		}
	case "static":
		if ch.Static != nil {
			result = append(result, collectNestedFormActions(ch.Static.Children, seen)...)
		}
	case "each":
		if ch.Each != nil {
			result = append(result, collectNestedFormActions(ch.Each.Children, seen)...)
		}
	}
	return result
}
