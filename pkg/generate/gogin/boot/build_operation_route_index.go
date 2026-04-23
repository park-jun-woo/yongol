//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=dos-guard
//ff:what buildOperationRouteIndex — OpenAPI 의 operationId → "METHOD /path" 매핑

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildOperationRouteIndex walks the OpenAPI doc and maps each
// operationId to its gin route key ("METHOD /path"). Returns empty map
// when the doc is nil so callers can still iterate safely.
func buildOperationRouteIndex(fs *yongol.Fullstack) map[string]string {
	idx := map[string]string{}
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return idx
	}
	for path, pi := range fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pi.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			idx[op.OperationID] = method + " " + openAPIPathToGin(path)
		}
	}
	return idx
}
