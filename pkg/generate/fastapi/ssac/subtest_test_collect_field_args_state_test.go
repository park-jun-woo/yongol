//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsState — State 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsState(t *testing.T) {

	op := ir.Op{Kind: ir.OpState, State: &ir.StateOp{Inputs: []ir.FieldArg{faKey("s")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "s" {
		t.Errorf("got %v", got)
	}

}
