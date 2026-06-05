//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectOpIDs — OpenAPI 문서의 모든 operationId 집합 수집 (컴포넌트 소비 필터/커버리지 판정용)

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectOpIDs returns the set of all non-empty operationIds declared across the
// OpenAPI document's path items. It is used both to filter component api.<Op>(
// candidates and as the universe for coverage rules (XMO-10/12).
func collectOpIDs(doc *openapi3.T) map[string]struct{} {
	out := make(map[string]struct{})
	if doc == nil || doc.Paths == nil {
		return out
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range operationsOf(item) {
			if op != nil && op.OperationID != "" {
				out[op.OperationID] = struct{}{}
			}
		}
	}
	return out
}
