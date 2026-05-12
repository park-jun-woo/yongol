//ff:func feature=generate type=util control=iteration dimension=1
//ff:what ChildNode 트리를 순회하여 폼 필드가 있는 중첩 액션을 수집한다
package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectNestedFormActions walks the ChildNode tree for nested actions with fields.
func collectNestedFormActions(nodes []stmlparser.ChildNode, seen map[string]bool) []actionEntry {
	var result []actionEntry
	for _, ch := range nodes {
		result = appendChildNodeFormActions(result, ch, seen)
	}
	return result
}
