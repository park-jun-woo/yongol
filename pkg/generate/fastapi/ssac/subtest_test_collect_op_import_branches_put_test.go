//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectOpImportBranchesPut — Put 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectOpImportBranchesPut(t *testing.T) {

	d := newImportData()
	collectOpImport(&d, ir.Op{Kind: ir.OpPut, Put: &ir.PutOp{Model: "Item"}}, "f")
	if !d.UsesUpdate || !d.Models["Item"] {
		t.Errorf("got %+v", d)
	}

}
