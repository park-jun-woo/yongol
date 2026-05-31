//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsDeleteNil — DeleteNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsDeleteNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpDelete}) != nil {
		t.Error("expected nil")
	}

}
