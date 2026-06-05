//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what objectProp — 테스트용 object(맵, additionalProperties: string) schema ref 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// objectProp creates an object(map) schema ref with string additionalProperties.
func objectProp() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		AdditionalProperties: openapi3.AdditionalProperties{
			Schema: stringProp(),
		},
	}}
}
