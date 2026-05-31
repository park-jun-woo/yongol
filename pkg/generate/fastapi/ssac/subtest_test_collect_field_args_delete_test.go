//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsDelete — Delete 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsDelete(t *testing.T) {

	op := ir.Op{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Args: []ir.FieldArg{faKey("d")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "d" {
		t.Errorf("got %v", got)
	}

}
