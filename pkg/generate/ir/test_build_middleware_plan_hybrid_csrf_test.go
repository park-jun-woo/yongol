//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildMiddlewarePlanHybridCSRF -- hybrid 인증 → CSRF 커스텀 설정 + HybridBearerSkip 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildMiddlewarePlanHybridCSRF(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/hybrid",
				Auth: &manifest.Auth{
					Claims: map[string]manifest.ClaimDef{"UserID": {Key: "uid"}},
					Mode:   "hybrid",
					Csrf: &manifest.CsrfConfig{
						Enabled:    true,
						CookieName: "MY-CSRF",
						MaxAge:     3600,
					},
				},
			},
		},
	}
	ps := &prepared.State{
		Auth: prepared.Auth{
			Present:      true,
			Mode:         "hybrid",
			CsrfRequired: true,
			Raw:          fs.Manifest.Backend.Auth,
		},
	}

	plan := BuildMiddlewarePlan(fs, ps)

	if plan.CSRF == nil {
		t.Fatal("CSRF should be non-nil for hybrid auth")
	}
	if plan.CSRF.CookieName != "MY-CSRF" {
		t.Errorf("CSRF.CookieName = %q, want %q", plan.CSRF.CookieName, "MY-CSRF")
	}
	if plan.CSRF.MaxAge != 3600 {
		t.Errorf("CSRF.MaxAge = %d, want 3600", plan.CSRF.MaxAge)
	}
	if !plan.CSRF.HybridBearerSkip {
		t.Error("HybridBearerSkip should be true for hybrid mode")
	}
}
