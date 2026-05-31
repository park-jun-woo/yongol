//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestConvertGetBranches_ZeroCov(t *testing.T) {
	// pagination args separated, list wrapper, nil result
	seq := ssac.Sequence{
		Type:   ssac.SeqGet,
		Model:  "Course.List",
		Inputs: map[string]string{"PerPage": "request.per_page", "InstructorID": "request.iid"},
		Result: &ssac.Result{Var: "courses", Type: "Course", Wrapper: "Page"},
	}
	op := convertGet(seq, &yongol.Fullstack{})
	if op.Kind != OpGet || op.Get == nil {
		t.Fatalf("expected OpGet, got %+v", op)
	}
	if len(op.Get.PaginationArgs) != 1 || op.Get.PaginationArgs[0].Key != "PerPage" {
		t.Errorf("pagination args = %+v", op.Get.PaginationArgs)
	}
	if len(op.Get.Args) != 1 || op.Get.Args[0].Key != "InstructorID" {
		t.Errorf("where args = %+v", op.Get.Args)
	}
	if !op.Get.IsList {
		t.Error("expected IsList for Page wrapper")
	}

	// []-prefixed list type
	op = convertGet(ssac.Sequence{Model: "Course.List", Result: &ssac.Result{Var: "cs", Type: "[]Course"}}, nil)
	if !op.Get.IsList {
		t.Error("expected IsList for []-prefixed type")
	}

	// count result type
	op = convertGet(ssac.Sequence{Model: "Course.Count", Result: &ssac.Result{Var: "n", Type: "int64"}}, nil)
	if !op.Get.IsCount {
		t.Error("expected IsCount for int64")
	}

	// nil result
	op = convertGet(ssac.Sequence{Model: "Course.X"}, nil)
	if op.Get.VarName != "" {
		t.Errorf("expected empty VarName for nil result, got %q", op.Get.VarName)
	}
}

func TestIsCountResultType_ZeroCov(t *testing.T) {
	for _, ty := range []string{"int64", "int32", "int", "uint64"} {
		if !isCountResultType(ty) {
			t.Errorf("%q should be count", ty)
		}
	}
	if isCountResultType("Course") {
		t.Error("Course should not be count")
	}
}

func TestConvertPost_ZeroCov2(t *testing.T) {
	op := convertPost(ssac.Sequence{Model: "Course.Create", Inputs: map[string]string{"Title": "request.title"}, Result: &ssac.Result{Var: "c", Type: "Course", Wrapper: "Page"}})
	if op.Kind != OpPost || op.Post == nil {
		t.Fatalf("expected OpPost, got %+v", op)
	}
	if op.Post.Model != "Course" || op.Post.Method != "Create" {
		t.Errorf("model/method = %q/%q", op.Post.Model, op.Post.Method)
	}
	if !op.Post.IsList {
		t.Error("expected IsList for wrapper")
	}
	// nil result branch
	op = convertPost(ssac.Sequence{Model: "Course.Create"})
	if op.Post.VarName != "" {
		t.Errorf("expected empty VarName, got %q", op.Post.VarName)
	}
}

func TestConvertCall_ZeroCov2(t *testing.T) {
	// default ErrStatus, with result
	op := convertCall(ssac.Sequence{Model: "Mail.Send", Inputs: map[string]string{"To": "x"}, Result: &ssac.Result{Var: "r", Type: "Result"}})
	if op.Kind != OpCall || op.Call == nil {
		t.Fatalf("expected OpCall, got %+v", op)
	}
	if op.Call.Package != "Mail" || op.Call.TargetFeature != "mail" || op.Call.Function != "Send" {
		t.Errorf("call meta = %+v", op.Call)
	}
	if op.Call.ErrStatus != 500 {
		t.Errorf("default ErrStatus = %d, want 500", op.Call.ErrStatus)
	}
	if op.Call.ResultVar != "r" {
		t.Errorf("ResultVar = %q", op.Call.ResultVar)
	}
	// custom ErrStatus, no result
	op = convertCall(ssac.Sequence{Model: "Mail.Send", ErrStatus: 422, Message: "boom"})
	if op.Call.ErrStatus != 422 || op.Call.Message != "boom" {
		t.Errorf("custom call meta = %+v", op.Call)
	}
}

func TestConvertEval_ZeroCov2(t *testing.T) {
	op := convertEval(ssac.Sequence{Model: "Rule.Check", Message: "no"})
	if op.Kind != OpEval || op.Eval == nil {
		t.Fatalf("expected OpEval, got %+v", op)
	}
	if op.Eval.StatusCode != 400 {
		t.Errorf("default status = %d, want 400", op.Eval.StatusCode)
	}
	op = convertEval(ssac.Sequence{Model: "Rule.Check", ErrStatus: 409})
	if op.Eval.StatusCode != 409 {
		t.Errorf("custom status = %d", op.Eval.StatusCode)
	}
}

func TestConvertResponse_ZeroCov(t *testing.T) {
	// fields branch (sorted)
	op := convertResponse(ssac.Sequence{Fields: map[string]string{"b": "y", "a": "x"}})
	if op.Kind != OpResponse || op.Response == nil {
		t.Fatalf("expected OpResponse, got %+v", op)
	}
	if len(op.Response.Fields) != 2 || op.Response.Fields[0].Name != "a" {
		t.Errorf("fields not sorted: %+v", op.Response.Fields)
	}
	// target branch
	op = convertResponse(ssac.Sequence{Target: "course"})
	if op.Response.SingleVar != "course" {
		t.Errorf("SingleVar = %q", op.Response.SingleVar)
	}
	// empty branch
	op = convertResponse(ssac.Sequence{})
	if op.Response.SingleVar != "" || len(op.Response.Fields) != 0 {
		t.Errorf("expected empty response, got %+v", op.Response)
	}
}

func TestConvertState_ZeroCov2(t *testing.T) {
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{
			{ID: "reservation", Symbol: "Reservation", Transitions: []statemachine.Transition{
				{From: "pending", To: "cancelled", Event: "cancel"},
				{From: "active", To: "cancelled", Event: "cancel"},
			}},
		},
	}
	op := convertState(ssac.Sequence{DiagramID: "reservation", Transition: "cancel"}, fs)
	if op.Kind != OpState || op.State == nil {
		t.Fatalf("expected OpState, got %+v", op)
	}
	if op.State.StatusCode != 409 {
		t.Errorf("default status = %d", op.State.StatusCode)
	}
	if len(op.State.AllowedFromStates) != 2 {
		t.Errorf("allowed from states = %+v", op.State.AllowedFromStates)
	}
	// custom status, nil fs
	op = convertState(ssac.Sequence{DiagramID: "x", Transition: "y", ErrStatus: 423}, nil)
	if op.State.StatusCode != 423 || len(op.State.AllowedFromStates) != 0 {
		t.Errorf("nil-fs state = %+v", op.State)
	}
}

func TestConvertAuth_ZeroCov2(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{Ownerships: []rego.OwnershipMapping{{Resource: "project", Table: "projects", Column: "owner_id"}}},
		},
		DDLTables: []ddl.Table{{Name: "projects", PrimaryKey: []string{"id"}}},
	}
	// ownership populated
	op := convertAuth(ssac.Sequence{Action: "delete", Resource: "project", Inputs: map[string]string{"ResourceID": "project.ID"}}, fs)
	if op.Kind != OpAuth || op.Auth == nil {
		t.Fatalf("expected OpAuth, got %+v", op)
	}
	if op.Auth.StatusCode != 403 {
		t.Errorf("default status = %d", op.Auth.StatusCode)
	}
	if op.Auth.Ownership == nil || op.Auth.Ownership.Table != "projects" || op.Auth.Ownership.ResourcePK != "id" {
		t.Errorf("ownership = %+v", op.Auth.Ownership)
	}
	// zero ResourceID → no ownership; custom status
	op = convertAuth(ssac.Sequence{Action: "read", Resource: "project", ErrStatus: 401, Inputs: map[string]string{"ResourceID": "0"}}, fs)
	if op.Auth.Ownership != nil || op.Auth.StatusCode != 401 {
		t.Errorf("expected no ownership, status 401, got %+v / %d", op.Auth.Ownership, op.Auth.StatusCode)
	}
	// no matching policy resource
	op = convertAuth(ssac.Sequence{Resource: "other", Inputs: map[string]string{"ResourceID": "x.ID"}}, fs)
	if op.Auth.Ownership != nil {
		t.Errorf("expected no ownership for unmatched resource")
	}
	// nil fs
	op = convertAuth(ssac.Sequence{Resource: "project", Inputs: map[string]string{"ResourceID": "x.ID"}}, nil)
	if op.Auth.Ownership != nil {
		t.Errorf("expected no ownership for nil fs")
	}
}

func TestIsResourceIDZeroIR_ZeroCov(t *testing.T) {
	for _, z := range []string{"", "  ", "0", `""`, "''", "nil", "NULL", "Null"} {
		if !isResourceIDZeroIR(z) {
			t.Errorf("%q should be zero", z)
		}
	}
	for _, nz := range []string{"project.ID", "5", "x"} {
		if isResourceIDZeroIR(nz) {
			t.Errorf("%q should be non-zero", nz)
		}
	}
}

func TestFindTablePK_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{
		{Name: "projects", PrimaryKey: []string{"id"}},
		{Name: "logs"},
	}}
	if got := findTablePK(fs, "PROJECTS"); got != "id" {
		t.Errorf("findTablePK = %q, want id", got)
	}
	if got := findTablePK(fs, "logs"); got != "" {
		t.Errorf("findTablePK no PK = %q, want empty", got)
	}
	if got := findTablePK(fs, "missing"); got != "" {
		t.Errorf("findTablePK missing = %q, want empty", got)
	}
}

func TestConvertVerifyPassword_ZeroCov2(t *testing.T) {
	op := convertVerifyPassword(ssac.Sequence{
		Model: "User", EmailCol: "email", EmailExpr: "req.Email",
		HashCol: "password_hash", PasswordExpr: "req.Password",
		ErrStatus: 401, Message: "bad",
		Result: &ssac.Result{Var: "u", Type: "User"},
	})
	if op.Kind != OpVerifyPassword || op.VerifyPW == nil {
		t.Fatalf("expected OpVerifyPassword, got %+v", op)
	}
	if op.VerifyPW.ResultVar != "u" || op.VerifyPW.ResultType != "User" {
		t.Errorf("verify pw result = %+v", op.VerifyPW)
	}
	// nil result
	op = convertVerifyPassword(ssac.Sequence{Model: "User"})
	if op.VerifyPW.ResultVar != "" {
		t.Errorf("expected empty ResultVar, got %q", op.VerifyPW.ResultVar)
	}
}

func TestConvertSequence_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}
	cases := []struct {
		seqType string
		seq     ssac.Sequence
		want    OpKind
	}{
		{ssac.SeqGet, ssac.Sequence{Type: ssac.SeqGet, Model: "A.B"}, OpGet},
		{ssac.SeqPost, ssac.Sequence{Type: ssac.SeqPost, Model: "A.B"}, OpPost},
		{ssac.SeqPut, ssac.Sequence{Type: ssac.SeqPut, Model: "A.B"}, OpPut},
		{ssac.SeqDelete, ssac.Sequence{Type: ssac.SeqDelete, Model: "A.B"}, OpDelete},
		{ssac.SeqEmpty, ssac.Sequence{Type: ssac.SeqEmpty, Target: "x"}, OpEmpty},
		{ssac.SeqExists, ssac.Sequence{Type: ssac.SeqExists, Target: "x"}, OpExists},
		{ssac.SeqAuth, ssac.Sequence{Type: ssac.SeqAuth, Resource: "p"}, OpAuth},
		{ssac.SeqState, ssac.Sequence{Type: ssac.SeqState, DiagramID: "d", Transition: "t"}, OpState},
		{ssac.SeqCall, ssac.Sequence{Type: ssac.SeqCall, Model: "A.B"}, OpCall},
		{ssac.SeqEval, ssac.Sequence{Type: ssac.SeqEval, Model: "A.B"}, OpEval},
		{ssac.SeqPublish, ssac.Sequence{Type: ssac.SeqPublish, Topic: "t"}, OpPublish},
		{ssac.SeqVerifyPassword, ssac.Sequence{Type: ssac.SeqVerifyPassword, Model: "User"}, OpVerifyPassword},
		{ssac.SeqResponse, ssac.Sequence{Type: ssac.SeqResponse, Target: "x"}, OpResponse},
	}
	for _, c := range cases {
		op, err := convertSequence(c.seq, fs)
		if err != nil {
			t.Errorf("%s: unexpected error %v", c.seqType, err)
			continue
		}
		if op.Kind != c.want {
			t.Errorf("%s: Kind = %d, want %d", c.seqType, op.Kind, c.want)
		}
	}
	// unknown type → error
	if _, err := convertSequence(ssac.Sequence{Type: "bogus"}, fs); err == nil {
		t.Error("expected error for unknown sequence type")
	}
}
