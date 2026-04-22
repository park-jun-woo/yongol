//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-301 긍정/부정 케이스 — default-src 에 '*' / 'unsafe-eval' 포함 시 WARNING

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
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

func TestSec301_NoConfig(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{},
		},
	}
	if diags := sec301CspPermissive(fs); len(diags) != 0 {
		t.Fatalf("missing block should not fire, got %+v", diags)
	}
}
