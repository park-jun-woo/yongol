//ff:func feature=validate type=test control=sequence topic=ssac-func
//ff:what zz_zerocov_test — ssac_func 0% 규칙 (Run / xfs42 / xfs43 / xfs45 / xsf46) 단위 테스트
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_ZeroCov(t *testing.T) {
	if diags := Run(&yongol.Fullstack{}); len(diags) != 0 {
		t.Fatalf("empty fullstack → 0 diags, got %d: %+v", len(diags), diags)
	}
}
