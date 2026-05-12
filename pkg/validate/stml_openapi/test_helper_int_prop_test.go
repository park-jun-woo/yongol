//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what intProp — 테스트용 integer schema ref 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// intProp creates a simple integer schema ref.
func intProp() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}}
}
