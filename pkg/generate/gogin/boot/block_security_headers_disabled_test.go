//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what blockSecurityHeaders — middleware.SecurityHeadersMiddleware 등록 + Config 조립
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockSecurityHeaders_Disabled(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{SecurityHeaders: &pmanifest.SecurityHeadersConfig{Enabled: &disabled}},
	}}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	if block.Active == nil || block.Active(fs) {
		t.Errorf("disabled security headers must carry an Active predicate returning false")
	}
	if len(block.Lines) != 0 {
		t.Errorf("disabled block must emit no lines, got %v", block.Lines)
	}
}
