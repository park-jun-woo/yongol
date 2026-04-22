//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-302 긍정/부정 케이스 — HSTS max_age < 180일 이면 WARNING

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec302_ShortMaxAge(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					HSTS: &pmanifest.HSTSConfig{MaxAge: 86400}, // 1 day
				},
			},
		},
	}
	diags := sec302HSTSShort(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 SEC-302 WARNING, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "SEC-302") {
		t.Errorf("missing SEC-302 in message: %q", diags[0].Message)
	}
}

func TestSec302_OneYearPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					HSTS: &pmanifest.HSTSConfig{MaxAge: 31536000},
				},
			},
		},
	}
	if diags := sec302HSTSShort(fs); len(diags) != 0 {
		t.Fatalf("1 year max-age should pass, got %+v", diags)
	}
}

func TestSec302_PreloadMinimumPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					HSTS: &pmanifest.HSTSConfig{MaxAge: 15552000}, // 180d
				},
			},
		},
	}
	if diags := sec302HSTSShort(fs); len(diags) != 0 {
		t.Fatalf("preload minimum should pass, got %+v", diags)
	}
}

func TestSec302_ZeroMaxAgeSkipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					HSTS: &pmanifest.HSTSConfig{MaxAge: 0},
				},
			},
		},
	}
	// max_age=0 is explicit disable, not misconfiguration → no warning.
	if diags := sec302HSTSShort(fs); len(diags) != 0 {
		t.Fatalf("max_age=0 should not fire, got %+v", diags)
	}
}

func TestSec302_NoHSTSBlock(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{},
			},
		},
	}
	if diags := sec302HSTSShort(fs); len(diags) != 0 {
		t.Fatalf("missing HSTS block should not fire, got %+v", diags)
	}
}
