//ff:func feature=gen-ir type=test control=sequence
//ff:what TestServicePlanOpenAPIMeta -- ServicePlan OpenAPI 메타데이터 이식 검증 (SuccessStatus/PathParams/QueryParams/BodyFields)
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestServicePlanSuccessStatus(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/courses", &openapi3.PathItem{
				Post: &openapi3.Operation{
					OperationID: "CreateCourse",
					Responses: openapi3.NewResponses(
						openapi3.WithStatus(201, &openapi3.ResponseRef{
							Value: openapi3.NewResponse().WithDescription("Created"),
						}),
					),
				},
			}),
		),
	}

	sf := &ssac.ServiceFunc{
		Name:     "CreateCourse",
		FileName: "create_course.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqPost, Model: "Course.Create", Result: &ssac.Result{Var: "c", Type: "Course"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{OpenAPIDoc: doc})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}
	if plan.SuccessStatus != 201 {
		t.Errorf("SuccessStatus = %d, want 201", plan.SuccessStatus)
	}
	if plan.HTTPMethod != "POST" {
		t.Errorf("HTTPMethod = %q, want POST", plan.HTTPMethod)
	}
	if plan.URLPath != "/courses" {
		t.Errorf("URLPath = %q, want /courses", plan.URLPath)
	}
}
