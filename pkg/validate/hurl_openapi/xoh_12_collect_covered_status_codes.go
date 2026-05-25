//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what collectCoveredStatusCodes — hurl entries 에서 operationId 별 커버된 status code 집합 수집

package hurl_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

// collectCoveredStatusCodes returns operationId -> {statusCode -> true} from
// all hurl entries matched against the given OpenAPI routes.
func collectCoveredStatusCodes(entries []hurl.HurlEntry, routes []apiRoute) map[string]map[string]bool {
	covered := map[string]map[string]bool{}
	for _, e := range entries {
		if e.StatusCode == "" {
			continue
		}
		segs := normalizeHurlPath(e.Path)
		r := findExactRoute(segs, e.Method, routes)
		if r == nil || r.Op == nil || r.Op.OperationID == "" {
			continue
		}
		opID := r.Op.OperationID
		if covered[opID] == nil {
			covered[opID] = map[string]bool{}
		}
		covered[opID][e.StatusCode] = true
	}
	return covered
}
