//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
