//ff:func feature=validate type=test control=sequence topic=ssac-rego
//ff:what zz_zerocov_test — ssac_rego.Run 0% 커버리지 단위 테스트
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXsp29RegoAllowToSSaC_ZeroCov(t *testing.T) {
	// Empty fs → no rego pairs → no diags, but body executes.
	if diags := xsp29RegoAllowToSSaC(&yongol.Fullstack{}); len(diags) != 0 {
		t.Errorf("empty fs should yield 0 diags, got %+v", diags)
	}
	// nil fs → nil.
	if got := xsp29RegoAllowToSSaC(nil); got != nil {
		t.Errorf("nil fs should return nil, got %v", got)
	}
}
