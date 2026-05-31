//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-301 테스트 — 'self' 만 있는 strict 설정은 통과

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec301_StrictPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					CSP: &pmanifest.CSPConfig{
						Directives: map[string][]string{
							"default-src": {"'self'"},
						},
					},
				},
			},
		},
	}
	if diags := sec301CspPermissive(fs); len(diags) != 0 {
		t.Fatalf("strict default-src should pass, got %+v", diags)
	}
}
