//ff:func feature=generate type=util control=iteration dimension=1
//ff:what STML 페이지에서 폼 필드가 있는 모든 액션의 operationId와 필드명을 수집한다
package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectFormActionOps iterates all pages and returns actions that have fields.
func collectFormActionOps(pages []stmlparser.PageSpec) []actionEntry {
	seen := make(map[string]bool)
	var result []actionEntry
	for _, page := range pages {
		result = appendPageFormActions(result, page, seen)
	}
	return result
}
