//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsPost — Post 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsPost(t *testing.T) {

	op := ir.Op{Kind: ir.OpPost, Post: &ir.PostOp{Args: []ir.FieldArg{faKey("p")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "p" {
		t.Errorf("got %v", got)
	}

}
