//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what mergeDocPaths — 단일 도메인 doc 의 모든 path 항목을 merged doc 에 union (paths 는 XDO-90 하 전역 유일)

package yongol

import "github.com/getkin/kin-openapi/openapi3"

// mergeDocPaths unions every path item of doc into merged. Paths are globally
// unique across domains under XDO-90, so the union is collision-free. nil doc or
// nil Paths is a no-op.
func mergeDocPaths(merged, doc *openapi3.T) {
	if doc == nil || doc.Paths == nil {
		return
	}
	for path, item := range doc.Paths.Map() {
		merged.Paths.Set(path, item)
	}
}
