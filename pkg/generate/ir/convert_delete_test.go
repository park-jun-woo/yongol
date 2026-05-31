//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertDelete(t *testing.T) {
	op := convertDelete(ssac.Sequence{Model: "Reservation.Delete", Inputs: map[string]string{"ID": "request.ID"}})
	if op.Kind != OpDelete || op.Delete == nil {
		t.Fatalf("op = %+v", op)
	}
	if op.Delete.Model != "Reservation" || op.Delete.Method != "Delete" {
		t.Errorf("model/method = %q/%q", op.Delete.Model, op.Delete.Method)
	}
	if len(op.Delete.Args) != 1 {
		t.Errorf("args = %+v", op.Delete.Args)
	}
}
