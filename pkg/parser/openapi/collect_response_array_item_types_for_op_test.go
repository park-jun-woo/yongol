//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectResponseArrayItemTypesForOp — 무ID/무응답/비2xx/비JSON 스킵과 2xx JSON 수집 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectResponseArrayItemTypesForOp(t *testing.T) {
	result := map[string]map[string]map[string]string{}

	// missing operationId / nil responses → skipped
	collectResponseArrayItemTypesForOp(result, &openapi3.Operation{})
	collectResponseArrayItemTypesForOp(result, &openapi3.Operation{OperationID: "NoResp"})
	if len(result) != 0 {
		t.Fatalf("expected empty result, got %v", result)
	}

	arraySchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"photos": {Value: &openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type: &openapi3.Types{"object"},
					Properties: openapi3.Schemas{
						"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
					},
				}},
			}},
		},
	}

	// non-2xx and non-JSON contents are skipped
	op := &openapi3.Operation{OperationID: "GetUnit", Responses: &openapi3.Responses{}}
	op.Responses.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.NewContentWithJSONSchema(arraySchema),
	}})
	op.Responses.Set("204", &openapi3.ResponseRef{Value: &openapi3.Response{}})
	collectResponseArrayItemTypesForOp(result, op)
	if len(result) != 0 {
		t.Fatalf("non-2xx/no-content must be skipped, got %v", result)
	}

	// 2xx JSON response collects the types
	op.Responses.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.NewContentWithJSONSchema(arraySchema),
	}})
	collectResponseArrayItemTypesForOp(result, op)
	if result["GetUnit"]["photos"]["id"] != "integer" {
		t.Errorf("expected photos.id=integer, got %v", result)
	}
}
