//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingDuplicateGet(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "DuplicateVar",
		FileName: "dup.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:   ssac.SeqGet,
				Model:  "Course.FindByID",
				Inputs: map[string]string{"ID": "request.id"},
				Result: &ssac.Result{Var: "item", Type: "Course"},
			},
			{
				Type:   ssac.SeqGet,
				Model:  "Course.FindBySlug",
				Inputs: map[string]string{"Slug": "request.slug"},
				Result: &ssac.Result{Var: "item", Type: "Course"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if plan.Ops[0].Get.VarName != "item" {
		t.Errorf("Ops[0].Get.VarName = %q, want %q", plan.Ops[0].Get.VarName, "item")
	}
	if plan.Ops[1].Get.VarName != "item_result" {
		t.Errorf("Ops[1].Get.VarName = %q, want %q", plan.Ops[1].Get.VarName, "item_result")
	}
}

func TestVariableShadowingTriple(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "TripleVar",
		FileName: "triple.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "A.X", Result: &ssac.Result{Var: "v", Type: "A"}},
			{Type: ssac.SeqGet, Model: "B.Y", Result: &ssac.Result{Var: "v", Type: "B"}},
			{Type: ssac.SeqGet, Model: "C.Z", Result: &ssac.Result{Var: "v", Type: "C"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	names := []string{
		plan.Ops[0].Get.VarName,
		plan.Ops[1].Get.VarName,
		plan.Ops[2].Get.VarName,
	}
	want := []string{"v", "v_result", "v_result_result"}
	for i, n := range names {
		if n != want[i] {
			t.Errorf("Ops[%d].Get.VarName = %q, want %q", i, n, want[i])
		}
	}
}

func TestVariableShadowingNoDuplicate(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "NoDup",
		FileName: "nodup.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "A.X", Result: &ssac.Result{Var: "a", Type: "A"}},
			{Type: ssac.SeqGet, Model: "B.Y", Result: &ssac.Result{Var: "b", Type: "B"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if plan.Ops[0].Get.VarName != "a" {
		t.Errorf("VarName[0] = %q, want a", plan.Ops[0].Get.VarName)
	}
	if plan.Ops[1].Get.VarName != "b" {
		t.Errorf("VarName[1] = %q, want b", plan.Ops[1].Get.VarName)
	}
}

func TestVariableShadowingCallAndPost(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "MixedVars",
		FileName: "mixed.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqPost, Model: "A.Create",
				Result: &ssac.Result{Var: "result", Type: "A"}},
			{Type: ssac.SeqCall, Model: "pkg.Fn",
				Result: &ssac.Result{Var: "result", Type: "B"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if plan.Ops[0].Post.VarName != "result" {
		t.Errorf("Post.VarName = %q, want result", plan.Ops[0].Post.VarName)
	}
	if plan.Ops[1].Call.ResultVar != "result_result" {
		t.Errorf("Call.ResultVar = %q, want result_result", plan.Ops[1].Call.ResultVar)
	}
}

// TestVariableShadowingReservedUserCallSource verifies that when "user" is a
// reserved name, a VerifyPassword result "user" is renamed to "user_result"
// and subsequent @call FieldArg.Source references are updated from "user" to
// "user_result".
func TestVariableShadowingReservedUserCallSource(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Login",
		FileName: "login.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:         ssac.SeqVerifyPassword,
				Model:        "User",
				EmailCol:     "Email",
				EmailExpr:    "request.Email",
				HashCol:      "PasswordHash",
				PasswordExpr: "request.Password",
				Result:       &ssac.Result{Var: "user", Type: "User"},
			},
			{
				Type:   ssac.SeqCall,
				Model:  "auth.IssueToken",
				Inputs: map[string]string{"Email": "user.Email", "Role": "user.Role"},
				Result: &ssac.Result{Var: "token", Type: "TokenPair"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	// VerifyPassword result "user" should be renamed to "user_result"
	// because "user" is a reserved name.
	vpOp := plan.Ops[0]
	if vpOp.Kind != OpVerifyPassword {
		t.Fatalf("Ops[0].Kind = %d, want OpVerifyPassword", vpOp.Kind)
	}
	if vpOp.VerifyPW.ResultVar != "user_result" {
		t.Errorf("VerifyPW.ResultVar = %q, want %q", vpOp.VerifyPW.ResultVar, "user_result")
	}

	// @call args should have Source updated from "user" to "user_result".
	callOp := plan.Ops[1]
	if callOp.Kind != OpCall {
		t.Fatalf("Ops[1].Kind = %d, want OpCall", callOp.Kind)
	}
	emailArg := findArgByKey(callOp.Call.Args, "Email")
	if emailArg == nil {
		t.Fatal("missing Email arg")
	}
	if emailArg.Source != "user_result" {
		t.Errorf("Email.Source = %q, want %q", emailArg.Source, "user_result")
	}
	roleArg := findArgByKey(callOp.Call.Args, "Role")
	if roleArg == nil {
		t.Fatal("missing Role arg")
	}
	if roleArg.Source != "user_result" {
		t.Errorf("Role.Source = %q, want %q", roleArg.Source, "user_result")
	}
}

// TestVariableShadowingResponseSourceUpdate verifies that ResponseField.Source
// dot-notation references are updated after variable renaming.
func TestVariableShadowingResponseSourceUpdate(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Login",
		FileName: "login.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:         ssac.SeqVerifyPassword,
				Model:        "User",
				EmailCol:     "Email",
				EmailExpr:    "request.Email",
				HashCol:      "PasswordHash",
				PasswordExpr: "request.Password",
				Result:       &ssac.Result{Var: "user", Type: "User"},
			},
			{
				Type:   ssac.SeqCall,
				Model:  "auth.IssueToken",
				Inputs: map[string]string{"Email": "user.Email"},
				Result: &ssac.Result{Var: "token", Type: "TokenPair"},
			},
			{
				Type: ssac.SeqResponse,
				Fields: map[string]string{
					"access_token":  "token.AccessToken",
					"refresh_token": "token.RefreshToken",
					"email":         "user.Email",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	// Response op should have "user" references updated to "user_result".
	respOp := plan.Ops[2]
	if respOp.Kind != OpResponse {
		t.Fatalf("Ops[2].Kind = %d, want OpResponse", respOp.Kind)
	}

	for _, f := range respOp.Response.Fields {
		switch f.Name {
		case "email":
			if f.Source != "user_result.Email" {
				t.Errorf("email.Source = %q, want %q", f.Source, "user_result.Email")
			}
		case "access_token":
			// token is not renamed, so Source stays as-is.
			if f.Source != "token.AccessToken" {
				t.Errorf("access_token.Source = %q, want %q", f.Source, "token.AccessToken")
			}
		case "refresh_token":
			if f.Source != "token.RefreshToken" {
				t.Errorf("refresh_token.Source = %q, want %q", f.Source, "token.RefreshToken")
			}
		}
	}
}
