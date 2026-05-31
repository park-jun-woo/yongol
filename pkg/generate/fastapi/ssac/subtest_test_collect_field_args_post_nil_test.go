//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsPostNil — PostNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsPostNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpPost}) != nil {
		t.Error("expected nil")
	}

}
