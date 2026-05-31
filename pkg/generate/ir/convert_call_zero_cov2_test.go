//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
