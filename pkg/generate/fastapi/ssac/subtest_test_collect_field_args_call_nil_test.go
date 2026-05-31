//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsCallNil — CallNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsCallNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpCall}) != nil {
		t.Error("expected nil")
	}

}
