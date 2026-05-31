//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesEvalExternal — EvalExternal 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesEvalExternal(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "other", Function: "Calc"}}, "f")
	if !d.ExtPkgs["other"]["Calc"] {
		t.Errorf("got %+v", d.ExtPkgs)
	}

}
