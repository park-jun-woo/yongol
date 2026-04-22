//ff:func feature=gen-gogin type=test-helper control=iteration dimension=1
//ff:what buildDoc — 테스트용 최소 OpenAPI 문서 조립 (opSpec 목록 + 루트 security 토글)
package boot

import "github.com/getkin/kin-openapi/openapi3"

// buildDoc constructs a minimal *openapi3.T from a list of opSpec and a
// globalSecurity flag. Used exclusively by boot package tests.
func buildDoc(ops []opSpec, globalSecurity bool) *openapi3.T {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	if globalSecurity {
		doc.Security = openapi3.SecurityRequirements{{"bearerAuth": []string{}}}
	}
	for _, o := range ops {
		attachOp(doc, o)
	}
	return doc
}
