//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesCallSelfSkipped — CallSelfSkipped 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesCallSelfSkipped(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "f", Function: "Do"}}, "f")
	if len(d.ExtPkgs) != 0 {
		t.Errorf("expected self-import skipped, got %+v", d.ExtPkgs)
	}

}
