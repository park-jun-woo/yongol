//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestBuildModelFetchMap — fetch 응답 top-level prop → operationEntry 맵 구성 및 unknown op 스킵 검증
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildModelFetchMap(t *testing.T) {
	resp := &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Properties: openapi3.Schemas{
					"workflow": {Value: &openapi3.Schema{Type: &openapi3.Types{"object"}}},
					"items":    {Value: &openapi3.Schema{Type: &openapi3.Types{"array"}}},
				}}},
			},
		},
	}}
	op := &openapi3.Operation{Responses: openapi3.NewResponses(openapi3.WithStatus(200, resp))}
	entry := operationEntry{method: "get", op: op}
	opMap := map[string]operationEntry{"GetWorkflow": entry}

	fetches := []stml.FetchBlock{
		{OperationID: "GetWorkflow"},
		{OperationID: "Unknown"}, // skipped: not in opMap.
	}

	out := buildModelFetchMap(fetches, opMap)

	if len(out) != 2 {
		t.Fatalf("expected 2 props mapped, got %d: %+v", len(out), out)
	}
	if out["workflow"].op != op {
		t.Errorf("workflow should map to GetWorkflow entry, got %+v", out["workflow"])
	}
	if out["items"].op != op {
		t.Errorf("items should map to GetWorkflow entry, got %+v", out["items"])
	}
	if _, ok := out["Unknown"]; ok {
		t.Errorf("unknown operationId must not contribute props")
	}
}
