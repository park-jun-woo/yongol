//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesDeleteNil — DeleteNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesDeleteNil(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpDelete}, "f")
	if !d.UsesDelete || len(d.Models) != 0 {
		t.Errorf("got %+v", d)
	}

}
