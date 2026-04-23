//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_DisabledExplicitly — Enabled=false 시 inert block

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBlockSecurityHeaders_DisabledExplicitly ensures Enabled=false yields
// an inert block with an Active guard that keeps it out of main.go.
func TestBlockSecurityHeaders_DisabledExplicitly(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					Enabled: &disabled,
				},
			},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	if block.Active == nil || block.Active(fs) {
		t.Fatalf("disabled block must not be active")
	}
	if len(block.Lines) != 0 {
		t.Fatalf("disabled block should emit no lines, got %+v", block.Lines)
	}
}
