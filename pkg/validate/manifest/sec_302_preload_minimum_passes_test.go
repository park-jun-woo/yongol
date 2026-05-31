//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-302 테스트 — preload 최소 180일 max_age 는 통과

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
