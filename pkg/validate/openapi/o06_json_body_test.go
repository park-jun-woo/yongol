//ff:func feature=validate type=test-helper control=sequence topic=openapi-structural
//ff:what o06JSONBody — 스키마 ref 를 application/json content 로 감싼다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// o06JSONBody wraps a schema ref into application/json content. Used by O-6
// request/response unit tests.
func o06JSONBody(ref *openapi3.SchemaRef) openapi3.Content {
	return openapi3.Content{"application/json": &openapi3.MediaType{Schema: ref}}
}
