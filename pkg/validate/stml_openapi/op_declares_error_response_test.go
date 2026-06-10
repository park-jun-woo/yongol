//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what opDeclaresErrorResponse — nil op / nil responses / 2xx only / 4xx / 5xx / 와일드카드 판정 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOpDeclaresErrorResponse(t *testing.T) {
	resp := func() *openapi3.ResponseRef {
		desc := "r"
		return &openapi3.ResponseRef{Value: &openapi3.Response{Description: &desc}}
	}
	if opDeclaresErrorResponse(nil) {
		t.Errorf("nil op: want false")
	}
	if opDeclaresErrorResponse(&openapi3.Operation{}) {
		t.Errorf("nil responses: want false")
	}
	twoxx := &openapi3.Operation{Responses: openapi3.NewResponses(openapi3.WithStatus(200, resp()))}
	if opDeclaresErrorResponse(twoxx) {
		t.Errorf("2xx only: want false")
	}
	fourxx := &openapi3.Operation{Responses: openapi3.NewResponses(openapi3.WithStatus(200, resp()), openapi3.WithStatus(404, resp()))}
	if !opDeclaresErrorResponse(fourxx) {
		t.Errorf("404 declared: want true")
	}
	fivexx := &openapi3.Operation{Responses: openapi3.NewResponses(openapi3.WithStatus(500, resp()))}
	if !opDeclaresErrorResponse(fivexx) {
		t.Errorf("500 declared: want true")
	}
	wild := &openapi3.Operation{Responses: openapi3.NewResponses(openapi3.WithStatus(200, resp()))}
	wild.Responses.Set("5XX", resp())
	if !opDeclaresErrorResponse(wild) {
		t.Errorf("5XX wildcard: want true")
	}
	deflt := &openapi3.Operation{Responses: openapi3.NewResponses(openapi3.WithStatus(200, resp()))}
	deflt.Responses.Set("default", resp())
	if opDeclaresErrorResponse(deflt) {
		t.Errorf("default only: want false")
	}
}
