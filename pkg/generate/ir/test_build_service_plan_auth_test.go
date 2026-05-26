//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanAuth -- @auth 시퀀스 IR 변환 (action/resource/inputs/status 검증)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanAuth(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:     ssac.SeqAuth,
				Action:   "ArchiveWorkflow",
				Resource: "workflow",
				Inputs: map[string]string{
					"ResourceID": "wf.ID",
				},
				Message:   "Forbidden",
				ErrStatus: 403,
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Ops) != 1 {
		t.Fatalf("len(Ops) = %d, want 1", len(plan.Ops))
	}

	authOp := plan.Ops[0]
	if authOp.Kind != OpAuth {
		t.Fatalf("Ops[0].Kind = %d, want OpAuth", authOp.Kind)
	}
	if authOp.Auth.Action != "ArchiveWorkflow" {
		t.Errorf("Auth.Action = %q, want %q", authOp.Auth.Action, "ArchiveWorkflow")
	}
	if authOp.Auth.Resource != "workflow" {
		t.Errorf("Auth.Resource = %q, want %q", authOp.Auth.Resource, "workflow")
	}
	if authOp.Auth.StatusCode != 403 {
		t.Errorf("Auth.StatusCode = %d, want 403", authOp.Auth.StatusCode)
	}
	if authOp.Auth.Message != "Forbidden" {
		t.Errorf("Auth.Message = %q, want %q", authOp.Auth.Message, "Forbidden")
	}
	ridArg := findArgByKey(authOp.Auth.Inputs, "ResourceID")
	if ridArg == nil {
		t.Fatal("Auth.Inputs missing ResourceID key")
	}
	if ridArg.Source != "wf" || ridArg.Field != "ID" {
		t.Errorf("ResourceID arg = {Source:%q Field:%q}, want {wf ID}", ridArg.Source, ridArg.Field)
	}
}
