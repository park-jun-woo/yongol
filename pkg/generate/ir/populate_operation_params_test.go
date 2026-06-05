//ff:func feature=gen-ir type=test control=sequence
//ff:what TestPopulateOperationParams -- path/query 파라미터 분류·정렬 및 body 필드 정렬 수집 검증

package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestPopulateOperationParams(t *testing.T) {
	pathItem := &openapi3.PathItem{
		Parameters: openapi3.Parameters{
			{Value: &openapi3.Parameter{Name: "id", In: "path", Required: true}},
		},
	}
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			// query params intentionally out of alphabetical order
			{Value: &openapi3.Parameter{
				Name: "limit", In: "query",
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			}},
			{Value: &openapi3.Parameter{
				Name: "cursor", In: "query", Required: true,
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			}},
		},
		RequestBody: &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
				Type:     &openapi3.Types{"object"},
				Required: []string{"title"},
				Properties: openapi3.Schemas{
					"title": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
					"body":  {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			}),
		},
	}

	plan := &ServicePlan{}
	populateOperationParams(plan, pathItem, op, map[string]bool{}, map[string]bool{})

	if len(plan.PathParams) != 1 || plan.PathParams[0] != "id" {
		t.Errorf("PathParams = %v, want [id]", plan.PathParams)
	}
	// query params sorted by Name: cursor, limit
	if len(plan.QueryParams) != 2 {
		t.Fatalf("QueryParams len = %d, want 2", len(plan.QueryParams))
	}
	if plan.QueryParams[0].Name != "cursor" || plan.QueryParams[1].Name != "limit" {
		t.Errorf("QueryParams order = [%s %s], want [cursor limit]",
			plan.QueryParams[0].Name, plan.QueryParams[1].Name)
	}
	if plan.QueryParams[1].Type != "integer" {
		t.Errorf("limit Type = %q, want integer", plan.QueryParams[1].Type)
	}
	// body fields sorted by Name: body, title
	if len(plan.BodyFields) != 2 {
		t.Fatalf("BodyFields len = %d, want 2", len(plan.BodyFields))
	}
	if plan.BodyFields[0].Name != "body" || plan.BodyFields[1].Name != "title" {
		t.Errorf("BodyFields order = [%s %s], want [body title]",
			plan.BodyFields[0].Name, plan.BodyFields[1].Name)
	}
	if !plan.BodyFields[1].Required {
		t.Error("title should be Required")
	}
}
