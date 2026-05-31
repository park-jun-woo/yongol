//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesGetNil — GetNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesGetNil(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpGet}, "f")
	if !d.UsesSelect || len(d.Models) != 0 {
		t.Errorf("got %+v", d)
	}

}
