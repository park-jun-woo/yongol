//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsAuthNil — AuthNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsAuthNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpAuth}) != nil {
		t.Error("expected nil")
	}

}
