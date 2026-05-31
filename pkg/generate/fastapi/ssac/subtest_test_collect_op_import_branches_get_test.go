//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesGet — Get 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesGet(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpGet, Get: &ir.GetOp{Model: "User"}}, "f")
	if !d.UsesSelect || !d.Models["User"] {
		t.Errorf("got %+v", d)
	}

}
