//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-302 테스트 — HSTS 블록이 없으면 미발화

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
