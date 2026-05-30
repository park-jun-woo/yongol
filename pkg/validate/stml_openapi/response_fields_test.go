//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestResponseFields — 2xx 응답 스키마 property→type 맵 추출 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestResponseFields(t *testing.T) {
	// nil op → empty.
	if got := responseFields(nil); len(got) != 0 {
		t.Fatalf("nil op: expected empty, got %+v", got)
	}
	// nil responses → empty.
	if got := responseFields(&openapi3.Operation{}); len(got) != 0 {
		t.Fatalf("nil responses: expected empty, got %+v", got)
	}

	// No 200/201 → loop exhausted, empty.
	emptyOp := &openapi3.Operation{Responses: openapi3.NewResponses()}
	if got := responseFields(emptyOp); len(got) != 0 {
		t.Fatalf("no 2xx: expected empty, got %+v", got)
	}

	mkResp := func(props openapi3.Schemas) *openapi3.ResponseRef {
		return &openapi3.ResponseRef{Value: &openapi3.Response{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Properties: props}},
				},
			},
		}}
	}

	// Only 201 present → first iteration (200) continues, 201 yields props.
	op201 := &openapi3.Operation{
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(201, mkResp(openapi3.Schemas{
				"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			})),
		),
	}
	out := responseFields(op201)
	if out["id"].typ != "integer" {
		t.Fatalf("expected id integer from 201, got %+v", out)
	}
}
