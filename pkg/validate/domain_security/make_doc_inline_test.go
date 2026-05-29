//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=domain-security
//ff:what makeDocInline — in-memory OpenAPI 문서 생성 (테스트용)
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// makeDocInline creates an openapi3.T using kin-openapi for in-memory tests.
func makeDocInline(paths map[string]*openapi3.PathItem) *openapi3.T {
	p := &openapi3.Paths{}
	for path, item := range paths {
		p.Set(path, item)
	}
	return &openapi3.T{Paths: p}
}
