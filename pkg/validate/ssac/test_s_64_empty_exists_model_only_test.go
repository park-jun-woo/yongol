//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what S-64 test — @empty/@exists Target Model 한정 (5 cases)

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS64EmptyExistsModelOnly(t *testing.T) {
	tt := []struct {
		name      string
		funcName  string
		seqs      []parsessac.Sequence
		modelSet  []string
		varTypes  map[string]string
		structs   []string
		wantDiag  int
		wantSubst string
	}{
		{
			name:     "model-OK: @empty wf where wf is DDL Workflow row",
			funcName: "ActivateWorkflow",
			seqs: []parsessac.Sequence{
				{Type: parsessac.SeqGet, Model: "Workflow.FindByID", Result: &parsessac.Result{Type: "Workflow", Var: "wf"}, Line: 10},
				{Type: parsessac.SeqEmpty, Target: "wf", Message: "not found", Line: 11},
			},
			modelSet: []string{"Workflow"},
			varTypes: map[string]string{"SSaC.var.ActivateWorkflow.wf": "Workflow"},
			wantDiag: 0,
		},
		{
			name:     "func-resp-OK: @exists r where r is FuncResponse struct",
			funcName: "DetectChange",
			seqs: []parsessac.Sequence{
				{Type: parsessac.SeqCall, Package: "billing", Model: "DetectChange", Result: &parsessac.Result{Type: "billing.DetectChangeResp", Var: "r"}, Line: 20},
				{Type: parsessac.SeqExists, Target: "r", Message: "duplicate", Line: 21},
			},
			structs:  []string{"DetectChangeResp"},
			varTypes: map[string]string{"SSaC.var.DetectChange.r": "billing.DetectChangeResp"},
			wantDiag: 0,
		},
		{
			name:     "scalar-int: @empty org.CreditsBalance is dotted-field → ERROR",
			funcName: "ActivateWorkflow",
			seqs: []parsessac.Sequence{
				{Type: parsessac.SeqGet, Model: "Org.FindByID", Result: &parsessac.Result{Type: "Org", Var: "org"}, Line: 30},
				{Type: parsessac.SeqEmpty, Target: "org.CreditsBalance", Message: "Insufficient credits", ErrStatus: 402, Line: 31},
			},
			modelSet:  []string{"Org"},
			varTypes:  map[string]string{"SSaC.var.ActivateWorkflow.org": "Org"},
			wantDiag:  1,
			wantSubst: "dotted-field",
		},
		{
			name:     "scalar-string: @empty bound to a string-typed var → ERROR",
			funcName: "RenameUser",
			seqs: []parsessac.Sequence{
				{Type: parsessac.SeqEmpty, Target: "email", Message: "missing email", Line: 41},
			},
			varTypes:  map[string]string{"SSaC.var.RenameUser.email": "string"},
			wantDiag:  1,
			wantSubst: "must be a Model var",
		},
		{
			name:     "dotted-field-on-model: @empty wf.ID → ERROR even when wf is a model",
			funcName: "ActivateWorkflow",
			seqs: []parsessac.Sequence{
				{Type: parsessac.SeqGet, Model: "Workflow.FindByID", Result: &parsessac.Result{Type: "Workflow", Var: "wf"}, Line: 50},
				{Type: parsessac.SeqEmpty, Target: "wf.ID", Message: "missing id", Line: 51},
			},
			modelSet:  []string{"Workflow"},
			varTypes:  map[string]string{"SSaC.var.ActivateWorkflow.wf": "Workflow"},
			wantDiag:  1,
			wantSubst: "dotted-field",
		},
	}

	for _, tc := range tt {
		fs := buildS64Fixture(tc.funcName, tc.seqs, tc.modelSet, tc.varTypes, tc.structs)
		diags := s64EmptyExistsModelOnly(fs)
		if len(diags) != tc.wantDiag {
			t.Errorf("[%s] diag count: got %d want %d (%+v)", tc.name, len(diags), tc.wantDiag, diags)
			continue
		}
		if tc.wantDiag > 0 && tc.wantSubst != "" && !strings.Contains(diags[0].Message, tc.wantSubst) {
			t.Errorf("[%s] message: got %q want substring %q", tc.name, diags[0].Message, tc.wantSubst)
		}
	}
}
