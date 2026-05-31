//ff:func feature=validate type=test control=sequence topic=manifest-security-headers
//ff:what SEC-301 테스트 — security_headers 블록 미존재 시 미발화

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
