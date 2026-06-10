//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestBlockCsrf_BearerMode_RuntimeGated — bearer 빌드도 CSRF 블록을 방출 (런타임 authMode 게이트)

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBlockCsrf_BearerMode_RuntimeGated pins BUG-116 / Phase-B1: a
// manifest=bearer build with auth declared must still register the CSRF
// middleware. The block is no longer dormant for bearer — the emitted
// middleware no-ops at runtime in bearer mode (csrfRuntimeActive), but the
// registration must exist so a BACKEND_AUTH_MODE=cookie override on the same
// binary reaches a live CSRF check.
func TestBlockCsrf_BearerMode_RuntimeGated(t *testing.T) {
	raw := &pmanifest.Auth{Mode: "bearer"}
	a := prepared.Auth{Present: true, Mode: "bearer", Raw: raw}
	block := blockCsrf(a, "example.com/zenflow")
	if len(block.Lines) == 0 {
		t.Fatalf("bearer build with auth present must register CSRF (runtime-gated), got inert block")
	}
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "middleware.Csrf(middleware.CsrfConfig{") {
		t.Errorf("bearer block missing middleware.Csrf registration, got:\n%s", body)
	}
	// HybridBearerSkip stays false on a bearer build — the runtime gate, not
	// this static flag, governs bearer no-op behavior.
	if !strings.Contains(body, "HybridBearerSkip: false") {
		t.Errorf("bearer block should emit HybridBearerSkip: false, got:\n%s", body)
	}
}
