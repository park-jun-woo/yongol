//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsEval — Eval 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsEval(t *testing.T) {

	op := ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Args: []ir.FieldArg{faKey("e")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "e" {
		t.Errorf("got %v", got)
	}

}
