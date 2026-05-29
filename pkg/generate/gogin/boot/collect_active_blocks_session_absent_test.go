//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectActiveBlocks_SessionAbsent_BUG008 — session 비활성 시 session-init 블록 미포함

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCollectActiveBlocks_SessionAbsent_BUG008 pins the BUG-008
// regression: with no manifest.session block and no SSaC session.*
// calls, collectActiveBlocks must not panic and the result must not
// contain the session-init block.
func TestCollectActiveBlocks_SessionAbsent_BUG008(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/app"},
		},
	}
	p := prepared.New(fs)
	if p.ActiveBackends.Session != nil {
		t.Fatalf("prepared.State should leave Session nil; got %+v", p.ActiveBackends.Session)
	}
	blocks := collectActiveBlocks(fs, p, "example.com/app")
	for _, b := range blocks {
		if b.Name == "session-init" {
			t.Fatalf("session-init block should not appear when session is inactive")
		}
	}
}
