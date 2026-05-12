//ff:func feature=generate type=util control=iteration dimension=1
//ff:what ChildNode 트리를 순회하여 data-field가 없는 중첩 액션의 operationId를 수집한다

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectNestedFieldlessActions walks the ChildNode tree and adds operationIds
// of field-less actions to the result map.
func collectNestedFieldlessActions(nodes []stmlparser.ChildNode, result map[string]bool) {
	for _, ch := range nodes {
		appendFieldlessFromChild(ch, result)
	}
}
