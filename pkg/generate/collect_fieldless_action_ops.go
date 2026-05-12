//ff:func feature=generate type=util control=iteration dimension=1
//ff:what STML 페이지에서 data-field가 없는 모든 액션의 operationId를 수집한다

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectFieldlessActionOps returns the set of operationIds for actions that
// have no data-field descendants. These are body-less mutations from the STML
// perspective, even if OpenAPI defines a requestBody.
func collectFieldlessActionOps(pages []stmlparser.PageSpec) map[string]bool {
	result := make(map[string]bool)
	for _, page := range pages {
		collectPageFieldlessActions(page, result)
	}
	return result
}
