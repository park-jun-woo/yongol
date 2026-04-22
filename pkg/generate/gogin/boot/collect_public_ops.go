//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what collectPublicOps — 인증 불필요(opt-out 포함) operationID 목록 수집
package boot

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// collectPublicOps walks the OpenAPI doc and returns operationIds that do not
// require authorization — either via explicit `security: []` opt-out or
// because no global security is declared. Result is sorted for deterministic
// code generation.
func collectPublicOps(doc *openapi3.T) []string {
	if doc == nil || doc.Paths == nil {
		return nil
	}
	var ops []string
	for _, pi := range doc.Paths.Map() {
		for _, op := range pi.Operations() {
			if op.OperationID == "" {
				continue
			}
			if !requiresAuth(op, doc) {
				ops = append(ops, op.OperationID)
			}
		}
	}
	sort.Strings(ops)
	return ops
}
