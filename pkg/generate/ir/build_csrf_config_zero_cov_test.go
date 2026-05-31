//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBuildCSRFConfig_ZeroCov(t *testing.T) {
	// not active → nil.
	if c := buildCSRFConfig(&prepared.State{}); c != nil {
		t.Errorf("inactive csrf → nil")
	}
	// active cookie mode, full overrides.
	ps := &prepared.State{Auth: prepared.Auth{
		Present: true, Mode: "cookie", CsrfRequired: true,
		Raw: &manifest.Auth{Csrf: &manifest.CsrfConfig{
			Enabled: true, CookieName: "CK", HeaderName: "HK", MaxAge: 7200,
			ExemptPaths: []string{"/x"},
		}},
	}}
	c := buildCSRFConfig(ps)
	if c == nil || c.CookieName != "CK" || c.HeaderName != "HK" || c.MaxAge != 7200 {
		t.Errorf("csrf overrides = %+v", c)
	}
	// active hybrid mode, no csrf overrides → defaults + HybridBearerSkip.
	ps2 := &prepared.State{Auth: prepared.Auth{
		Present: true, Mode: "hybrid", CsrfRequired: true,
		Raw: &manifest.Auth{},
	}}
	c2 := buildCSRFConfig(ps2)
	if c2 == nil || !c2.HybridBearerSkip || c2.CookieName != "XSRF-TOKEN" {
		t.Errorf("hybrid csrf defaults = %+v", c2)
	}
}
