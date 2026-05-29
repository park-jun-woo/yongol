//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildMiddlewarePlanCookieCSRF -- cookie 인증 → CSRF 기본 설정 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildMiddlewarePlanCookieCSRF(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/cookie",
				Auth: &manifest.Auth{
					Claims: map[string]manifest.ClaimDef{"UserID": {Key: "uid"}},
					Mode:   "cookie",
				},
			},
		},
	}
	ps := &prepared.State{
		Auth: prepared.Auth{
			Present:      true,
			Mode:         "cookie",
			CsrfRequired: true,
			Raw:          fs.Manifest.Backend.Auth,
		},
	}

	plan := BuildMiddlewarePlan(fs, ps)

	if plan.CSRF == nil {
		t.Fatal("CSRF should be non-nil for cookie auth")
	}
	if plan.CSRF.CookieName != "XSRF-TOKEN" {
		t.Errorf("CSRF.CookieName = %q, want %q", plan.CSRF.CookieName, "XSRF-TOKEN")
	}
	if plan.CSRF.HybridBearerSkip {
		t.Error("HybridBearerSkip should be false for cookie mode")
	}
}
