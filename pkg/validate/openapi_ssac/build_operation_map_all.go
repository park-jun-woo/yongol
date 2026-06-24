//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-openapi
//ff:what buildOperationMapAll — 전체 도메인(또는 단일 사이트) OpenAPI 문서를 합쳐 operationId → Operation 맵 생성

package openapi_ssac

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildOperationMapAll merges buildOperationMap over every OpenAPI document
// returned by fs.AllOpenAPIDocs() — the per-domain docs in domain mode, or the
// singular doc keyed "" in single-site mode. operationIds are globally unique
// under XDO-90, so the merge never clobbers an operation.
func buildOperationMapAll(fs *yongol.Fullstack) map[string]*openapi3.Operation {
	out := make(map[string]*openapi3.Operation)
	for _, doc := range fs.AllOpenAPIDocs() {
		for id, op := range buildOperationMap(doc) {
			out[id] = op
		}
	}
	return out
}
