//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_APIProfile — api profile 문자열 그대로 전달 확인

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
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
	body := strings.Join(blockSecurityHeaders(fs, "example.com/zenflow").Lines, "\n")
	if !strings.Contains(body, `"api"`) {
		t.Fatalf("api profile not emitted; got:\n%s", body)
	}
}
