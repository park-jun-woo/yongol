//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newCreatedResponses — 단일 201 Responses 객체 생성

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newCreatedResponses returns an openapi3.Responses with a single 201 entry.
func newCreatedResponses() *openapi3.Responses {
	desc := "created"
	r := openapi3.NewResponses()
	r.Set("201", &openapi3.ResponseRef{Value: &openapi3.Response{Description: &desc}})
	return r
}
