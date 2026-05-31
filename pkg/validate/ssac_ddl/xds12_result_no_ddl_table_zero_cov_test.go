//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what zz_zerocov_test — ssac_ddl 0% (Run / xds12ResultNoDDLTable / collectFuncResultDiags) 단위 테스트
package ssac_ddl

import (
	"testing"
)

func TestXds12ResultNoDDLTable_ZeroCov(t *testing.T) {
	diags := xds12ResultNoDDLTable(fsWithResultZeroCov())
	if len(diags) != 1 {
		t.Fatalf("expected 1 XDS-12 diag, got %d: %+v", len(diags), diags)
	}
}
