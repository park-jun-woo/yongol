//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what collectCoveredOps — smoke.hurl entries 에서 커버된 operationId 집합 수집

package hurl_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

func collectCoveredOps(entries []hurl.HurlEntry, routes []apiRoute) map[string]bool {
	covered := map[string]bool{}
	for _, e := range entries {
		segs := normalizeHurlPath(e.Path)
		r := findExactRoute(segs, e.Method, routes)
		if r == nil || r.Op == nil || r.Op.OperationID == "" {
			continue
		}
		covered[r.Op.OperationID] = true
	}
	return covered
}
