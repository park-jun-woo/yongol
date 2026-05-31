//ff:func feature=gen-ir type=test control=sequence
//ff:what TestServicePlanOpenAPIMeta -- ServicePlan OpenAPI 메타데이터 이식 검증 (SuccessStatus/PathParams/QueryParams/BodyFields)
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestServicePlanBodyFields(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/courses", &openapi3.PathItem{
				Post: &openapi3.Operation{
					OperationID: "CreateCourse",
					RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
						Content: openapi3.Content{
							"application/json": &openapi3.MediaType{
								Schema: openapi3.NewSchemaRef("", &openapi3.Schema{
									Type: &openapi3.Types{"object"},
									Properties: openapi3.Schemas{
										"title": openapi3.NewSchemaRef("", &openapi3.Schema{
											Type: &openapi3.Types{"string"},
										}),
										"email": openapi3.NewSchemaRef("", &openapi3.Schema{
											Type:   &openapi3.Types{"string"},
											Format: "email",
										}),
									},
									Required: []string{"title"},
								}),
							},
						},
					}},
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
			{Type: ssac.SeqPost, Model: "Course.Create",
				Inputs: map[string]string{"Title": "request.title"},
				Result: &ssac.Result{Var: "c", Type: "Course"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{OpenAPIDoc: doc})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if len(plan.BodyFields) != 2 {
		t.Fatalf("len(BodyFields) = %d, want 2", len(plan.BodyFields))
	}
	// Sorted by name.
	if plan.BodyFields[0].Name != "email" || plan.BodyFields[0].Format != "email" {
		t.Errorf("BodyFields[0] = {%q, format:%q}, want {email, email}",
			plan.BodyFields[0].Name, plan.BodyFields[0].Format)
	}
	if plan.BodyFields[1].Name != "title" || !plan.BodyFields[1].Required {
		t.Errorf("BodyFields[1] = {%q, required:%v}, want {title, true}",
			plan.BodyFields[1].Name, plan.BodyFields[1].Required)
	}
}
