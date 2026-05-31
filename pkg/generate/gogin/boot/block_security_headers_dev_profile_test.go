//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_DevProfile — dev profile 문자열 그대로 전달 확인

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockSecurityHeaders_DevProfile verifies the dev profile is wired
// through unchanged — runtime middleware decides what to do with it.
func TestBlockSecurityHeaders_DevProfile(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					Profile: "dev",
				},
			},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	combined := strings.Join(block.Lines, "\n") + "\n" + strings.Join(block.Funcs, "\n")
	if !strings.Contains(combined, `"dev"`) {
		t.Fatalf("dev profile not emitted; got:\n%s", combined)
	}
}
