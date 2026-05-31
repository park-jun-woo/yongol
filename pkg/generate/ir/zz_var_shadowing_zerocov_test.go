//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVarShadowingZeroCov — resolveVariableShadowing / applyRenames / resolveVar 전 분기 직접 커버

package ir

import "testing"

func TestResolveVar_ZeroCov(t *testing.T) {
	declared := map[string]bool{}
	if got := resolveVar("user", declared); got != "user" {
		t.Errorf("first use = %q, want user", got)
	}
	if got := resolveVar("user", declared); got != "user_result" {
		t.Errorf("collision = %q, want user_result", got)
	}
	if got := resolveVar("user", declared); got != "user_result_result" {
		t.Errorf("double collision = %q, want user_result_result", got)
	}
}

func TestResolveVariableShadowing_AllOpKinds_ZeroCov(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{VarName: "x"}},
		{Kind: OpPost, Post: &PostOp{VarName: "x"}}, // collides with get → x_result
		{Kind: OpCall, Call: &CallOp{ResultVar: "x"}},
		{Kind: OpVerifyPassword, VerifyPW: &VerifyPasswordOp{ResultVar: "x"}},
		// downstream op referencing the original "x" via FieldArg and response.
		{Kind: OpEval, Eval: &EvalOp{Args: []FieldArg{{Source: "x", Field: "ID"}}}},
		{Kind: OpResponse, Response: &ResponseOp{Fields: []ResponseField{
			{Name: "a", Source: "x"},
			{Name: "b", Source: "x.Email"},
		}}},
	}
	// "x" reserved so even the first declaration shadows → exercises rename path.
	resolveVariableShadowing(ops, "x")

	if ops[0].Get.VarName == "x" {
		t.Errorf("get var should have been renamed away from reserved x")
	}
	// downstream FieldArg.Source must have been rewritten to the renamed var.
	if ops[4].Eval.Args[0].Source == "x" {
		t.Errorf("downstream FieldArg.Source still references old x")
	}
	// dotted response source rewritten on the var part only.
	if ops[5].Response.Fields[1].Source == "x.Email" {
		t.Errorf("dotted response source not rewritten: %q", ops[5].Response.Fields[1].Source)
	}
}

func TestApplyRenames_NoRenames_ZeroCov(t *testing.T) {
	op := &Op{Kind: OpGet, Get: &GetOp{Args: []FieldArg{{Source: "x"}}}}
	applyRenames(op, map[string]string{}) // empty → early return, no mutation
	if op.Get.Args[0].Source != "x" {
		t.Errorf("empty renames should not mutate")
	}
}
