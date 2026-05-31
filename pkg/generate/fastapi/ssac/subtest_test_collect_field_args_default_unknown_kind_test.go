//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsDefaultUnknownKind — DefaultUnknownKind 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsDefaultUnknownKind(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpKind(9999)}) != nil {
		t.Error("expected nil for unknown kind")
	}

}
