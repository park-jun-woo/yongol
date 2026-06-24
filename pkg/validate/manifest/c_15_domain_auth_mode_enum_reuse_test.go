//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what C-15 재사용 테스트 — SEC-403 의 validAuthModes 가 허용하는 모든 값을 C-15 도 허용

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestC15ReusesSEC403Enum asserts C-15 shares SEC-403's validAuthModes map
// rather than maintaining a private copy: every key the shared enum accepts
// (cookie/bearer/hybrid/"") must pass C-15 unflagged.
func TestC15ReusesSEC403Enum(t *testing.T) {
	for mode, ok := range validAuthModes {
		if !ok {
			continue
		}
		diags := c15DomainAuthModeEnum(fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {AuthMode: mode},
			"admin":  {AuthMode: "bearer"},
		}))
		if len(diags) != 0 {
			t.Errorf("validAuthModes[%q]=true but C-15 flagged it: %+v", mode, diags)
		}
	}
}
