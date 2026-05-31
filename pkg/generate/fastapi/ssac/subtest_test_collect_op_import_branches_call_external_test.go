//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesCallExternal — CallExternal 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesCallExternal(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "other", Function: "Do"}}, "f")
	if !d.ExtPkgs["other"]["Do"] {
		t.Errorf("got %+v", d.ExtPkgs)
	}

}
