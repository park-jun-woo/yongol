//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-301 테스트 — default-src 에 '*' 포함 시 WARNING

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec301_Wildcard(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					CSP: &pmanifest.CSPConfig{
						Directives: map[string][]string{
							"default-src": {"*"},
						},
					},
				},
			},
		},
	}
	diags := sec301CspPermissive(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "SEC-301") {
		t.Errorf("missing SEC-301 in message: %q", diags[0].Message)
	}
}
