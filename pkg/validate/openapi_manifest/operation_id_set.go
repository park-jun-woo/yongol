//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what OpenAPI 문서 전체 operationId 집합 수집

package openapi_manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

// operationIDSet collects all operationIds declared in the OpenAPI doc.
func operationIDSet(fs *yongol.Fullstack) map[string]bool {
	s := map[string]bool{}
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return s
	}
	for _, pi := range fs.OpenAPIDoc.Paths.Map() {
		for _, op := range pi.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			s[op.OperationID] = true
		}
	}
	return s
}
