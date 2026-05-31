//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsEvalNil — EvalNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsEvalNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpEval}) != nil {
		t.Error("expected nil")
	}

}
