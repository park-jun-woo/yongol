//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesCallNilOrEmptyPackage — CallNilOrEmptyPackage 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesCallNilOrEmptyPackage(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpCall}, "f")
	collectOpImport(&d, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: ""}}, "f")
	if len(d.ExtPkgs) != 0 {
		t.Errorf("got %+v", d.ExtPkgs)
	}

}
