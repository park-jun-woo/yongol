//ff:func feature=gen-react type=test control=sequence
//ff:what resolveAPIClientPlan — 모드별(bearer/cookie/hybrid/jwt 무선언/csrf 설정) 계획 도출 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveAPIClientPlan(t *testing.T) {
	newFS := func(auth *manifest.Auth) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: auth}}}
	}

	// backend.auth absent -> plain client
	plan, err := resolveAPIClientPlan(nil)
	if err != nil || plan.bearer || plan.cookie {
		t.Errorf("nil fs = (%+v,%v), want plain plan", plan, err)
	}

	// bearer mode without frontend.auth -> bearer, no refresh (downgrade)
	plan, err = resolveAPIClientPlan(newFS(&manifest.Auth{Mode: "bearer"}))
	if err != nil || !plan.bearer || plan.cookie || plan.refresh != nil {
		t.Errorf("bearer = (%+v,%v), want bearer plan without refresh", plan, err)
	}

	// jwt type without mode -> bearer (prepared BUG-014 rule, not raw cookie default)
	plan, err = resolveAPIClientPlan(newFS(&manifest.Auth{Type: "jwt"}))
	if err != nil || !plan.bearer || plan.cookie {
		t.Errorf("jwt no-mode = (%+v,%v), want bearer plan", plan, err)
	}

	// cookie default (empty mode, no type) -> cookie + csrf with runtime default names
	plan, err = resolveAPIClientPlan(newFS(&manifest.Auth{}))
	if err != nil || plan.bearer || !plan.cookie || !plan.csrf {
		t.Fatalf("cookie = (%+v,%v), want cookie+csrf plan", plan, err)
	}
	if plan.csrfCookieName != "XSRF-TOKEN" || plan.csrfHeaderName != "X-XSRF-TOKEN" {
		t.Errorf("csrf names = (%q,%q), want runtime defaults", plan.csrfCookieName, plan.csrfHeaderName)
	}

	// hybrid -> cookie path too
	plan, err = resolveAPIClientPlan(newFS(&manifest.Auth{Mode: "hybrid"}))
	if err != nil || !plan.cookie || !plan.csrf {
		t.Errorf("hybrid = (%+v,%v), want cookie+csrf plan", plan, err)
	}

	// manifest csrf overrides win
	plan, _ = resolveAPIClientPlan(newFS(&manifest.Auth{Mode: "cookie", Csrf: &manifest.CsrfConfig{Enabled: true, CookieName: "MY-COOKIE", HeaderName: "X-My-Header"}}))
	if plan.csrfCookieName != "MY-COOKIE" || plan.csrfHeaderName != "X-My-Header" {
		t.Errorf("csrf overrides = (%q,%q)", plan.csrfCookieName, plan.csrfHeaderName)
	}

	// csrf explicitly disabled -> cookie client without the csrf middleware
	plan, _ = resolveAPIClientPlan(newFS(&manifest.Auth{Mode: "cookie", Csrf: &manifest.CsrfConfig{Enabled: false}}))
	if !plan.cookie || plan.csrf {
		t.Errorf("csrf disabled = %+v, want cookie without csrf", plan)
	}
}
