//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what stringProp — 테스트용 string schema ref 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// stringProp creates a simple string schema ref.
func stringProp() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
}
