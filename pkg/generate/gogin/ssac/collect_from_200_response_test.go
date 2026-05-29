//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what collectFrom200Response 단위 테스트 (직접 body $ref + property $ref 수집)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func jsonResponseOp(method string, status int, schema *openapi3.SchemaRef) *openapi3.Operation {
	return &openapi3.Operation{
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(status, &openapi3.ResponseRef{
				Value: openapi3.NewResponse().
					WithDescription("ok").
					WithJSONSchemaRef(schema),
			}),
		),
	}
}

func TestCollectFrom200Response(t *testing.T) {
	t.Run("nil op is a no-op", func(t *testing.T) {
		out := map[string]bool{}
		collectFrom200Response(nil, "GET", out)
		if len(out) != 0 {
			t.Errorf("expected empty, got %v", out)
		}
	})

	t.Run("direct body ref", func(t *testing.T) {
		op := jsonResponseOp("GET", 200, &openapi3.SchemaRef{Ref: "#/components/schemas/Workflow"})
		out := map[string]bool{}
		collectFrom200Response(op, "GET", out)
		if !out["Workflow"] {
			t.Errorf("expected Workflow collected, got %v", out)
		}
	})

	t.Run("property refs collected on POST 201", func(t *testing.T) {
		schema := &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: openapi3.Schemas{
				"workflow": {Ref: "#/components/schemas/Workflow"},
				"action":   {Ref: "#/components/schemas/Action"},
				"count":    {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			},
		}}
		op := jsonResponseOp("POST", 201, schema)
		out := map[string]bool{}
		collectFrom200Response(op, "POST", out)
		if !out["Workflow"] || !out["Action"] {
			t.Errorf("expected Workflow+Action, got %v", out)
		}
		if len(out) != 2 {
			t.Errorf("scalar property should be ignored, got %v", out)
		}
	})
}
