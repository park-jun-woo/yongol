//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsCall — Call 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsCall(t *testing.T) {

	op := ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Args: []ir.FieldArg{faKey("c")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "c" {
		t.Errorf("got %v", got)
	}

}
