//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgLocation -- FieldArg.Location 분류 검증 (path/query/body/var/literal/user)
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgLocationBody(t *testing.T) {
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
			{
				Type:  ssac.SeqPost,
				Model: "Course.Create",
				Inputs: map[string]string{
					"Title":        "request.title",
					"InstructorID": "currentUser.ID",
				},
				Result: &ssac.Result{Var: "course", Type: "Course"},
			},
		},
	}

	fs := &yongol.Fullstack{OpenAPIDoc: doc}
	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	postOp := plan.Ops[0]
	if postOp.Kind != OpPost {
		t.Fatalf("Ops[0].Kind = %d, want OpPost", postOp.Kind)
	}

	titleArg := findArgByKey(postOp.Post.Args, "Title")
	if titleArg == nil {
		t.Fatal("missing Title arg")
	}
	if titleArg.Location != LocBody {
		t.Errorf("Title.Location = %q, want %q", titleArg.Location, LocBody)
	}

	instArg := findArgByKey(postOp.Post.Args, "InstructorID")
	if instArg == nil {
		t.Fatal("missing InstructorID arg")
	}
	if instArg.Location != LocUser {
		t.Errorf("InstructorID.Location = %q, want %q", instArg.Location, LocUser)
	}
}
