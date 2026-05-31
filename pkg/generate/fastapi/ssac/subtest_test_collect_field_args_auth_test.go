//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsAuth — Auth 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsAuth(t *testing.T) {

	op := ir.Op{Kind: ir.OpAuth, Auth: &ir.AuthOp{Inputs: []ir.FieldArg{faKey("au")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "au" {
		t.Errorf("got %v", got)
	}

}
