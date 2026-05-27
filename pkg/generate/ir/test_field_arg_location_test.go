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

func TestFieldArgLocationLiteral(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqPut,
				Model: "Workflow.UpdateStatus",
				Inputs: map[string]string{
					"Status": `"archived"`,
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	putOp := plan.Ops[0]
	statusArg := findArgByKey(putOp.Put.Args, "Status")
	if statusArg == nil {
		t.Fatal("missing Status arg")
	}
	if statusArg.Location != LocLiteral {
		t.Errorf("Status.Location = %q, want %q", statusArg.Location, LocLiteral)
	}
}

func TestFieldArgLocationVar(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Workflow.FindByID",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{Var: "wf", Type: "Workflow"},
			},
			{
				Type:  ssac.SeqPut,
				Model: "Workflow.UpdateStatus",
				Inputs: map[string]string{
					"ID": "wf.ID",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	putOp := plan.Ops[1]
	idArg := findArgByKey(putOp.Put.Args, "ID")
	if idArg == nil {
		t.Fatal("missing ID arg")
	}
	if idArg.Location != LocVar {
		t.Errorf("ID.Location = %q, want %q", idArg.Location, LocVar)
	}
}
