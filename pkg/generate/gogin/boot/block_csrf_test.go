//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what blockCsrf — middleware.Csrf 등록 (쿠키 인증 조건부, Phase005)

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_Inactive(t *testing.T) {
	// Default bearer (no csrf) → inert block with always-inactive predicate.
	block := blockCsrf(prepared.Auth{Present: false}, "example.com/zenflow")
	if block.Active == nil || block.Active(nil) {
		t.Errorf("non-csrf block must carry csrfAlwaysInactive predicate")
	}
	if len(block.Lines) != 0 {
		t.Errorf("inactive csrf must emit no lines, got %v", block.Lines)
	}
}

func TestBlockCsrf_HybridSetsBearerSkip(t *testing.T) {
	raw := &pmanifest.Auth{Mode: "hybrid", Csrf: &pmanifest.CsrfConfig{Enabled: true}}
	a := prepared.Auth{Present: true, Mode: "hybrid", CsrfRequired: true, Raw: raw}
	block := blockCsrf(a, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "HybridBearerSkip: true") {
		t.Errorf("hybrid mode must set HybridBearerSkip: true, got:\n%s", body)
	}
	if !strings.Contains(body, "middleware.Csrf(middleware.CsrfConfig{") {
		t.Errorf("must register Csrf middleware, got:\n%s", body)
	}
}
