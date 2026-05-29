//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-302 테스트 — max_age=0 은 명시적 비활성화로 간주 (미발화)

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
