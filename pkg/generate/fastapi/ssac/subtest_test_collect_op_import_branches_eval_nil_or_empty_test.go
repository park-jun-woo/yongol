//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesEvalNilOrEmpty — EvalNilOrEmpty 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesEvalNilOrEmpty(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpEval}, "f")
	if len(d.ExtPkgs) != 0 {
		t.Errorf("got %+v", d.ExtPkgs)
	}

}
