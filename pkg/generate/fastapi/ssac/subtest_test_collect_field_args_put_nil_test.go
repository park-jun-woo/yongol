//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsPutNil — PutNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsPutNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpPut}) != nil {
		t.Error("expected nil")
	}

}
