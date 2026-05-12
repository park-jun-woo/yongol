//ff:func feature=generate type=util control=iteration dimension=1
//ff:what 페이지의 액션과 중첩 자식에서 폼 필드가 있는 액션을 result에 추가한다
package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// appendPageFormActions appends form actions from a single page (top-level
// actions + nested children) to the result slice. It skips already-seen
// operationIds.
func appendPageFormActions(result []actionEntry, page stmlparser.PageSpec, seen map[string]bool) []actionEntry {
	for _, a := range page.Actions {
		if len(a.Fields) > 0 && !seen[a.OperationID] {
			seen[a.OperationID] = true
			result = append(result, toActionEntry(a))
		}
	}
	result = append(result, collectNestedFormActions(page.Children, seen)...)
	return result
}
