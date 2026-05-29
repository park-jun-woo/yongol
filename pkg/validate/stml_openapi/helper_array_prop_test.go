//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what arrayProp — 테스트용 array schema ref 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// arrayProp creates an array schema ref.
func arrayProp(itemType string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{itemType}}},
	}}
}
