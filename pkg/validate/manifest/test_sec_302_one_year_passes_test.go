//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-302 테스트 — 1년 max_age 는 통과

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
