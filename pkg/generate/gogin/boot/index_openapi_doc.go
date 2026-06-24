//ff:func feature=gen-gogin type=util control=iteration dimension=2 topic=dos-guard
//ff:what indexOpenAPIDoc — 한 OpenAPI 문서의 operationId → "METHOD <prefix><path>" 를 idx 에 병합

package boot

import "github.com/getkin/kin-openapi/openapi3"

// indexOpenAPIDoc walks one OpenAPI document and merges each operationId into
// idx as its gin route key ("METHOD <routePrefix><path>"). routePrefix is the
// domain route-group prefix the doc's relative paths are mounted under (empty
// for single-site, where openapi paths already match the runtime FullPath).
//
// operationIds are globally unique across domains (XDO-90), so merging several
// domain docs into one idx never collides.
//
// Nested loop structure (path → methods) is intrinsic to the OpenAPI document
// shape, so this func declares dimension=2 rather than flattening.
func indexOpenAPIDoc(idx map[string]string, doc *openapi3.T, routePrefix string) {
	if doc == nil || doc.Paths == nil {
		return
	}
	for path, pi := range doc.Paths.Map() {
		for method, op := range pi.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			idx[op.OperationID] = method + " " + routePrefix + openAPIPathToGin(path)
		}
	}
}
