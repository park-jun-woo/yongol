//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestCheckStateInputType_ZeroCov(t *testing.T) {
	g := &rule.Ground{Types: map[string]string{
		"SSaC.var.Fn.bad": "int64",
	}}
	fn := ssac.ServiceFunc{Name: "Fn", FileName: "f.ssac"}
	// "bad" resolves to int64 (non-string) → diag; "lit" is a quoted literal → string-compatible, no diag
	seq := ssac.Sequence{Type: "state", Line: 3, Inputs: map[string]string{"k": "bad"}}
	diags := checkStateInputType(g, fn, seq)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	// string-compatible input → no diag
	seq2 := ssac.Sequence{Type: "state", Line: 3, Inputs: map[string]string{"k": `"draft"`}}
	if d := checkStateInputType(g, fn, seq2); len(d) != 0 {
		t.Errorf("expected 0 diags for string literal, got %v", d)
	}
}
