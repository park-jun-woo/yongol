//ff:func feature=gen-ir type=test control=sequence
//ff:what TestServicePlanOpenAPIMeta -- ServicePlan OpenAPI 메타데이터 이식 검증 (SuccessStatus/PathParams/QueryParams/BodyFields)
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestServicePlanPathAndQueryParams(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/courses/{id}", &openapi3.PathItem{
				Parameters: openapi3.Parameters{
					&openapi3.ParameterRef{Value: &openapi3.Parameter{
						Name: "id", In: "path", Required: true,
						Schema: openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{"integer"}}),
					}},
				},
				Get: &openapi3.Operation{
					OperationID: "GetCourse",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "cursor", In: "query",
							Schema: openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{"string"}}),
						}},
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "limit", In: "query", Required: true,
							Schema: openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{"integer"}}),
						}},
					},
					Responses: openapi3.NewResponses(
						openapi3.WithStatus(200, &openapi3.ResponseRef{
							Value: openapi3.NewResponse().WithDescription("OK"),
						}),
					),
				},
			}),
		),
	}

	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "Course.FindByID",
				Inputs: map[string]string{"ID": "request.id"},
				Result: &ssac.Result{Var: "c", Type: "Course"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{OpenAPIDoc: doc})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if len(plan.PathParams) != 1 || plan.PathParams[0] != "id" {
		t.Errorf("PathParams = %v, want [id]", plan.PathParams)
	}
	if len(plan.QueryParams) != 2 {
		t.Fatalf("len(QueryParams) = %d, want 2", len(plan.QueryParams))
	}
	// Sorted by name.
	if plan.QueryParams[0].Name != "cursor" {
		t.Errorf("QueryParams[0].Name = %q, want cursor", plan.QueryParams[0].Name)
	}
	if plan.QueryParams[1].Name != "limit" || !plan.QueryParams[1].Required {
		t.Errorf("QueryParams[1] = {%q, required:%v}, want {limit, true}",
			plan.QueryParams[1].Name, plan.QueryParams[1].Required)
	}
}
