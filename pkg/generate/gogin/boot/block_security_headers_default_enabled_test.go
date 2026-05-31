//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=security-headers
//ff:what blockSecurityHeaders — middleware.SecurityHeadersMiddleware 등록 + Config 조립
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockSecurityHeaders_DefaultEnabled(t *testing.T) {
	block := blockSecurityHeaders(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		`envBool("BACKEND_SECURITY_HEADERS_ENABLED", true)`,
		`envString("BACKEND_SECURITY_HEADERS_PROFILE", "production")`,
		"r.Use(middleware.SecurityHeadersMiddleware(buildSecurityHeadersCfg(",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("blockSecurityHeaders missing %q, got:\n%s", must, body)
		}
	}
	if len(block.Funcs) != 1 {
		t.Errorf("must emit buildSecurityHeadersCfg helper, got %d funcs", len(block.Funcs))
	}
}
