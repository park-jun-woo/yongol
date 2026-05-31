//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
