//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanCall -- @call 시퀀스 IR 변환 (pkg.Func + 결과 바인딩 + 기본 상태코드)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanCall(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Refresh",
		FileName: "refresh.ssac",
		Imports:  []string{"github.com/park-jun-woo/ssac/pkg/auth"},
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqCall,
				Model: "auth.RefreshRotate",
				Inputs: map[string]string{
					"RefreshToken": "request.refresh_token",
				},
				Result: &ssac.Result{Var: "rotated", Type: "auth.RefreshRotateResponse"},
			},
			{
				Type: ssac.SeqResponse,
				Fields: map[string]string{
					"access_token":  "rotated.AccessToken",
					"refresh_token": "rotated.RefreshToken",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Imports) != 1 {
		t.Fatalf("len(Imports) = %d, want 1", len(plan.Imports))
	}

	if len(plan.Ops) != 2 {
		t.Fatalf("len(Ops) = %d, want 2", len(plan.Ops))
	}

	callOp := plan.Ops[0]
	if callOp.Kind != OpCall {
		t.Fatalf("Ops[0].Kind = %d, want OpCall", callOp.Kind)
	}
	if callOp.Call.Package != "auth" {
		t.Errorf("Call.Package = %q, want %q", callOp.Call.Package, "auth")
	}
	if callOp.Call.Function != "RefreshRotate" {
		t.Errorf("Call.Function = %q, want %q", callOp.Call.Function, "RefreshRotate")
	}
	if callOp.Call.ResultVar != "rotated" {
		t.Errorf("Call.ResultVar = %q, want %q", callOp.Call.ResultVar, "rotated")
	}
	if callOp.Call.ErrStatus != 500 {
		t.Errorf("Call.ErrStatus = %d, want 500 (default)", callOp.Call.ErrStatus)
	}

	respOp := plan.Ops[1]
	if respOp.Kind != OpResponse {
		t.Fatalf("Ops[1].Kind = %d, want OpResponse", respOp.Kind)
	}
	if len(respOp.Response.Fields) != 2 {
		t.Fatalf("len(Response.Fields) = %d, want 2", len(respOp.Response.Fields))
	}
}
