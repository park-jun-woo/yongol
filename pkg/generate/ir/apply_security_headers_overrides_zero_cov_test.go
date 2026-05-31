//ff:func feature=gen-ir type=test control=sequence
//ff:what ir 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestApplySecurityHeadersOverrides_ZeroCov(t *testing.T) {
	cfg := &SecurityHeadersConfig{Profile: "production"}
	sh := &manifest.SecurityHeadersConfig{
		Profile:        "dev",
		HSTS:           &manifest.HSTSConfig{MaxAge: 100, IncludeSubDomains: true, Preload: true},
		CSP:            &manifest.CSPConfig{ReportOnly: true, Directives: map[string][]string{"default-src": {"'self'"}}},
		XFrameOptions:  "DENY",
		ReferrerPolicy: "no-referrer",
	}
	applySecurityHeadersOverrides(cfg, sh)
	if cfg.Profile != "dev" || cfg.HSTSMaxAge != 100 || cfg.XFrameOptions != "DENY" {
		t.Errorf("overrides not applied: %#v", cfg)
	}
	if !cfg.CSPReportOnly {
		t.Errorf("dev profile should force CSPReportOnly")
	}
}
