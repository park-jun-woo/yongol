//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-openapi
//ff:what buildOperationMethodMapAll — 전체 도메인(또는 단일 사이트) OpenAPI 문서를 합쳐 operationId → (method, Operation) 맵 생성

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildOperationMethodMapAll merges buildOperationMethodMap over every OpenAPI
// document returned by fs.AllOpenAPIDocs(). Like buildOperationMapAll it relies
// on XDO-90's global operationId uniqueness so per-domain entries never
// overwrite one another.
func buildOperationMethodMapAll(fs *yongol.Fullstack) map[string]OperationEntry {
	out := make(map[string]OperationEntry)
	for _, doc := range fs.AllOpenAPIDocs() {
		for id, entry := range buildOperationMethodMap(doc) {
			out[id] = entry
		}
	}
	return out
}
