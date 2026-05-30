//ff:func feature=validate type=test topic=design-manifest
//ff:what zz_zerocov_test — design_manifest.Run 0% 커버리지 단위 테스트
package design_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_ZeroCov(t *testing.T) {
	if diags := Run(&yongol.Fullstack{}); len(diags) != 0 {
		t.Fatalf("empty fullstack → 0 diags, got %d: %+v", len(diags), diags)
	}
}
