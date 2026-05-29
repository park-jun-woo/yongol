//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what makeDoc — 테스트용 OpenAPI doc fixture 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// makeDoc builds an OpenAPI doc with the given paths.
func makeDoc(paths map[string]*openapi3.PathItem) *openapi3.T {
	p := &openapi3.Paths{}
	for path, item := range paths {
		p.Set(path, item)
	}
	return &openapi3.T{Paths: p}
}
