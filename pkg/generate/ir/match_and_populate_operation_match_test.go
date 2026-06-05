//ff:func feature=gen-ir type=test control=sequence
//ff:what TestMatchAndPopulateOperationMatch — TestMatchAndPopulateOperation -- operationID 매칭 시 plan 채우고 true, 미스 시 false 검증

package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMatchAndPopulateOperationMatch(t *testing.T) {
	pathParam := &openapi3.Parameter{Name: "id", In: "path", Required: true}
	pathItem := &openapi3.PathItem{
		Delete: &openapi3.Operation{
			OperationID: "DeleteCourse",
			Parameters:  openapi3.Parameters{{Value: pathParam}},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(204, &openapi3.ResponseRef{
					Value: openapi3.NewResponse().WithDescription("No Content"),
				}),
			),
		},
	}

	plan := &ServicePlan{}
	pathParams := map[string]bool{"id": true}
	queryParams := map[string]bool{}

	ok := matchAndPopulateOperation("DeleteCourse", "/courses/{id}", pathItem, plan, pathParams, queryParams)
	if !ok {
		t.Fatal("expected match=true")
	}
	if plan.HTTPMethod != "DELETE" {
		t.Errorf("HTTPMethod = %q, want DELETE", plan.HTTPMethod)
	}
	if plan.URLPath != "/courses/{id}" {
		t.Errorf("URLPath = %q, want /courses/{id}", plan.URLPath)
	}
	if plan.SuccessStatus != 204 {
		t.Errorf("SuccessStatus = %d, want 204", plan.SuccessStatus)
	}
	if len(plan.PathParams) != 1 || plan.PathParams[0] != "id" {
		t.Errorf("PathParams = %v, want [id]", plan.PathParams)
	}
}
