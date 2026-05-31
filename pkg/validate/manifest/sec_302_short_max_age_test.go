//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-302 테스트 — max_age 가 180일 미만이면 WARNING

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
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
