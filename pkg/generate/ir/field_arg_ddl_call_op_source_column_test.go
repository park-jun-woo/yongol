//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgDDL -- FieldArg.ColumnName/IsPK DDL 매핑 + @call/@auth SourceColumn 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgDDLCallOpSourceColumn(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Login",
		FileName: "login.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:   ssac.SeqCall,
				Model:  "auth.IssueToken",
				Inputs: map[string]string{"Email": "user.Email"},
				Result: &ssac.Result{Var: "token", Type: "TokenPair"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	callOp := plan.Ops[0]
	if callOp.Kind != OpCall {
		t.Fatalf("Ops[0].Kind = %d, want OpCall", callOp.Kind)
	}

	emailArg := findArgByKey(callOp.Call.Args, "Email")
	if emailArg == nil {
		t.Fatal("missing Email arg")
	}
	if emailArg.SourceColumn != "email" {
		t.Errorf("Email.SourceColumn = %q, want %q", emailArg.SourceColumn, "email")
	}
	if emailArg.Source != "user" {
		t.Errorf("Email.Source = %q, want %q", emailArg.Source, "user")
	}
}
