//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-301 테스트 — 'unsafe-eval' 포함 시 WARNING

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec301_UnsafeEval(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					CSP: &pmanifest.CSPConfig{
						Directives: map[string][]string{
							"default-src": {"'self'", "'unsafe-eval'"},
						},
					},
				},
			},
		},
	}
	if len(sec301CspPermissive(fs)) != 1 {
		t.Fatal("expected SEC-301 WARNING for 'unsafe-eval'")
	}
}
