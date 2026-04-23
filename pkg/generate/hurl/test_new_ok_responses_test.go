//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newOKResponses — 단일 200 Responses 객체 생성

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newOKResponses returns an openapi3.Responses with a single 200 entry.
func newOKResponses() *openapi3.Responses {
	desc := "ok"
	r := openapi3.NewResponses()
	r.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{Description: &desc}})
	return r
}
