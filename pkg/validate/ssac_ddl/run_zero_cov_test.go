//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what zz_zerocov_test — ssac_ddl 0% (Run / xds12ResultNoDDLTable / collectFuncResultDiags) 단위 테스트
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_ZeroCov(t *testing.T) {
	if diags := Run(&yongol.Fullstack{}); len(diags) != 0 {
		t.Fatalf("empty fullstack → 0 diags, got %d: %+v", len(diags), diags)
	}
}
