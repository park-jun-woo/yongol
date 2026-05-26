//ff:func feature=validate type=test control=sequence topic=ssac-authz
//ff:what TestRun -- SSaC↔Authz Run 호출 성공 검증 (빈 입력 → 빈 결과)

package ssac_authz

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := Run(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags for empty fullstack, got %d", len(diags))
	}
}
