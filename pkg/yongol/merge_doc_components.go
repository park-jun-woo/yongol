//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what mergeDocComponents — 단일 도메인 doc 의 Schemas/SecuritySchemes 를 merged doc 에 union (이름 충돌 시 last-wins)

package yongol

import "github.com/getkin/kin-openapi/openapi3"

// mergeDocComponents unions doc's component Schemas and SecuritySchemes into
// merged. Names collide only on shared definitions (e.g. bearerAuth); last
// writer wins, which is harmless for the membership/existence checks the merged
// view feeds. nil doc or nil Components is a no-op.
func mergeDocComponents(merged, doc *openapi3.T) {
	if doc == nil || doc.Components == nil {
		return
	}
	for n, s := range doc.Components.Schemas {
		merged.Components.Schemas[n] = s
	}
	for n, s := range doc.Components.SecuritySchemes {
		merged.Components.SecuritySchemes[n] = s
	}
}
