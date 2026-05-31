//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgLocation -- FieldArg.Location 분류 검증 (path/query/body/var/literal/user)
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgLocation(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/courses/{id}", &openapi3.PathItem{
				Parameters: openapi3.Parameters{
					&openapi3.ParameterRef{Value: &openapi3.Parameter{
						Name: "id", In: "path",
					}},
				},
				Get: &openapi3.Operation{
					OperationID: "GetCourse",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "cursor", In: "query",
							Schema: openapi3.NewSchemaRef("", &openapi3.Schema{
								Type: &openapi3.Types{"string"},
							}),
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
		Feature:  "course",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.FindByID",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{Var: "course", Type: "Course"},
			},
			{
				Type:   ssac.SeqResponse,
				Target: "course",
			},
		},
	}

	fs := &yongol.Fullstack{OpenAPIDoc: doc}
	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0]
	if getOp.Kind != OpGet {
		t.Fatalf("Ops[0].Kind = %d, want OpGet", getOp.Kind)
	}
	idArg := findArgByKey(getOp.Get.Args, "ID")
	if idArg == nil {
		t.Fatal("missing ID arg")
	}
	if idArg.Location != LocPath {
		t.Errorf("ID.Location = %q, want %q", idArg.Location, LocPath)
	}
}
