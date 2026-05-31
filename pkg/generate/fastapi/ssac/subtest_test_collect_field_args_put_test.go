//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsPut — Put 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsPut(t *testing.T) {

	op := ir.Op{Kind: ir.OpPut, Put: &ir.PutOp{Args: []ir.FieldArg{faKey("u")}}}
	if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "u" {
		t.Errorf("got %v", got)
	}

}
