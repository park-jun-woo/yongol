//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesEvalSelfSkipped — EvalSelfSkipped 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesEvalSelfSkipped(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "f", Function: "Calc"}}, "f")
	if len(d.ExtPkgs) != 0 {
		t.Errorf("got %+v", d.ExtPkgs)
	}

}
