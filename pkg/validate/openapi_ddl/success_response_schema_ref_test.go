//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what successResponseSchemaRef — op 의 관용 2xx(GET=200) application/json 스키마 ref 추출, 없으면 nil

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSuccessResponseSchemaRef(t *testing.T) {
	if successResponseSchemaRef(nil, "GET") != nil {
		t.Error("nil op should return nil")
	}
	if successResponseSchemaRef(&openapi3.Operation{}, "GET") != nil {
		t.Error("nil Responses should return nil")
	}
	// GET 200 with JSON body → returns schema
	op := &openapi3.Operation{Responses: jsonResp("200", compRef("Rule", "id"))}
	if ref := successResponseSchemaRef(op, "GET"); ref == nil || ref.Ref != "#/components/schemas/Rule" {
		t.Errorf("expected Rule ref, got %+v", ref)
	}
	// GET with only a 404 (no conventional 2xx for GET=200) → nil
	op404 := &openapi3.Operation{Responses: jsonResp("404", inlineRef("error"))}
	if ref := successResponseSchemaRef(op404, "GET"); ref != nil {
		t.Errorf("expected nil for missing 200, got %+v", ref)
	}
	// 2xx present but no application/json content → nil
	r := openapi3.NewResponses()
	r.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{Content: openapi3.Content{
		"text/plain": &openapi3.MediaType{Schema: inlineRef("x")},
	}}})
	if ref := successResponseSchemaRef(&openapi3.Operation{Responses: r}, "GET"); ref != nil {
		t.Errorf("expected nil for non-json content, got %+v", ref)
	}
}
