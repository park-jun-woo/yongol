//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsGetNil — GetNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsGetNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpGet}) != nil {
		t.Error("expected nil for nil Get")
	}

}
