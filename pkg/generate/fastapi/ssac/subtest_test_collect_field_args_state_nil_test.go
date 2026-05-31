//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsStateNil — StateNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsStateNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpState}) != nil {
		t.Error("expected nil")
	}

}
