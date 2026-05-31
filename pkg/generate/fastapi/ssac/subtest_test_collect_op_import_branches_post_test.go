//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesPost — Post 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesPost(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpPost, Post: &ir.PostOp{Model: "Item"}}, "f")
	if !d.Models["Item"] {
		t.Errorf("got %+v", d)
	}

}
