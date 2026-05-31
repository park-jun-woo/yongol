//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesDelete — Delete 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesDelete(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Model: "Item"}}, "f")
	if !d.UsesDelete || !d.Models["Item"] {
		t.Errorf("got %+v", d)
	}

}
