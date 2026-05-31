//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsPublishNil — PublishNil 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsPublishNil(t *testing.T) {

	if collectFieldArgs(ir.Op{Kind: ir.OpPublish}) != nil {
		t.Error("expected nil")
	}

}
