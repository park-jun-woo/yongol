//ff:func feature=generate type=util control=iteration dimension=1
//ff:what 페이지의 top-level 액션과 중첩 자식에서 field-less 액션 operationId를 수집한다

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectPageFieldlessActions collects operationIds from field-less actions in
// a single page (top-level actions + nested children).
func collectPageFieldlessActions(page stmlparser.PageSpec, result map[string]bool) {
	for _, a := range page.Actions {
		if len(a.Fields) == 0 {
			result[a.OperationID] = true
		}
	}
	collectNestedFieldlessActions(page.Children, result)
}
