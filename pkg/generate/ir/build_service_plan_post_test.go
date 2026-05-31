//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanPost -- @post 시퀀스 IR 변환 + 트랜잭션 필요 판정 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanPost(t *testing.T) {
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
			{
				Type: ssac.SeqResponse,
				Fields: map[string]string{
					"course": "course",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if !plan.UsesTransaction {
		t.Error("UsesTransaction = false, want true for plan with @post")
	}
	if len(plan.Ops) != 2 {
		t.Fatalf("len(Ops) = %d, want 2", len(plan.Ops))
	}

	postOp := plan.Ops[0]
	if postOp.Kind != OpPost {
		t.Fatalf("Ops[0].Kind = %d, want OpPost", postOp.Kind)
	}
	if postOp.Post.Model != "Course" {
		t.Errorf("Post.Model = %q, want %q", postOp.Post.Model, "Course")
	}
	if postOp.Post.Method != "Create" {
		t.Errorf("Post.Method = %q, want %q", postOp.Post.Method, "Create")
	}
	if postOp.Post.VarName != "course" {
		t.Errorf("Post.VarName = %q, want %q", postOp.Post.VarName, "course")
	}
	if len(postOp.Post.Args) != 2 {
		t.Fatalf("len(Post.Args) = %d, want 2", len(postOp.Post.Args))
	}

	respOp := plan.Ops[1]
	if respOp.Kind != OpResponse {
		t.Fatalf("Ops[1].Kind = %d, want OpResponse", respOp.Kind)
	}
	if len(respOp.Response.Fields) != 1 {
		t.Fatalf("len(Response.Fields) = %d, want 1", len(respOp.Response.Fields))
	}
	if respOp.Response.Fields[0].Name != "course" || respOp.Response.Fields[0].Source != "course" {
		t.Errorf("Response.Fields[0] = {%q %q}, want {course course}", respOp.Response.Fields[0].Name, respOp.Response.Fields[0].Source)
	}
}
