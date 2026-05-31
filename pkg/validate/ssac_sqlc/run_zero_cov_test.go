//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_ZeroCov(t *testing.T) {
	// Empty Fullstack runs every sub-rule without panicking; diags empty.
	diags := Run(&yongol.Fullstack{})
	if len(diags) != 0 {
		t.Fatalf("empty fullstack should yield 0 diags, got %d: %+v", len(diags), diags)
	}
}
