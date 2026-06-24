//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what mergedOpenAPIDoc — 도메인별 OpenAPI doc 의 Paths/Components 를 정렬 순서로 union doc 에 합침

package yongol

import "github.com/getkin/kin-openapi/openapi3"

// mergedOpenAPIDoc builds the union document consumed by MergedOpenAPIView,
// merging each domain's doc (in sorted name order) into one. nil domain docs are
// tolerated by the per-aspect mergers.
func (fs *Fullstack) mergedOpenAPIDoc() *openapi3.T {
	merged := &openapi3.T{
		Paths: openapi3.NewPaths(),
		Components: &openapi3.Components{
			Schemas:         openapi3.Schemas{},
			SecuritySchemes: openapi3.SecuritySchemes{},
		},
	}
	for _, name := range fs.DomainNames() {
		doc := fs.DomainOpenAPIDocs[name]
		mergeDocPaths(merged, doc)
		mergeDocComponents(merged, doc)
	}
	return merged
}
