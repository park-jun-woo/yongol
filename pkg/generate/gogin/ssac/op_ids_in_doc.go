//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what opIDsInDoc — OpenAPI doc 의 모든 operationId 집합 반환 (도메인별 단일 방출 필터용)

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// opIDsInDoc returns the set of operationIds declared in doc. Domain-mode
// service-method generation runs once per domain over fs.DomainView(name); the
// shared internal/service package would otherwise re-emit every ServiceFunc's
// method for each domain. Skipping any ServiceFunc whose Name is absent from
// this set makes each operationId emit exactly once, by its owning domain.
// A nil doc / nil Paths yields an empty (non-nil) set.
func opIDsInDoc(doc *openapi3.T) map[string]bool {
	ids := make(map[string]bool)
	if doc == nil || doc.Paths == nil {
		return ids
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			ids[op.OperationID] = true
		}
	}
	return ids
}
