//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_APIProfile — api profile 문자열 그대로 전달 확인

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockSecurityHeaders_APIProfile — api profile just propagates the
// string, runtime drops CSP.
func TestBlockSecurityHeaders_APIProfile(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					Profile: "api",
				},
			},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	combined := strings.Join(block.Lines, "\n") + "\n" + strings.Join(block.Funcs, "\n")
	if !strings.Contains(combined, `"api"`) {
		t.Fatalf("api profile not emitted; got:\n%s", combined)
	}
}
