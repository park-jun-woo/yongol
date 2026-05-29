//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what resolveSecurityHeaders — manifest.security_headers + production 기본값 병합

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveSecurityHeaders_Defaults(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{nil, {}, {Manifest: &pmanifest.ProjectConfig{}}} {
		out := resolveSecurityHeaders(fs)
		if out.Profile != defaultSHProfile {
			t.Errorf("profile = %q, want %q", out.Profile, defaultSHProfile)
		}
		if out.HSTSMaxAge != defaultHSTSMaxAge || !out.CSPEnabled {
			t.Errorf("default preset wrong: %+v", out)
		}
		if out.XFrameOptions != defaultXFrameOptions {
			t.Errorf("x-frame = %q, want %q", out.XFrameOptions, defaultXFrameOptions)
		}
	}
}

func TestResolveSecurityHeaders_Overrides(t *testing.T) {
	cspOff := false
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{SecurityHeaders: &pmanifest.SecurityHeadersConfig{
			Profile:        "API",
			HSTS:           &pmanifest.HSTSConfig{MaxAge: 100, IncludeSubDomains: false, Preload: true},
			CSP:            &pmanifest.CSPConfig{Enabled: &cspOff, ReportOnly: true},
			XFrameOptions:  "SAMEORIGIN",
			ReferrerPolicy: "no-referrer",
		}},
	}}
	out := resolveSecurityHeaders(fs)
	if out.Profile != "api" {
		t.Errorf("profile should lowercase, got %q", out.Profile)
	}
	if out.HSTSMaxAge != 100 || out.HSTSIncludeSubs || !out.HSTSPreload {
		t.Errorf("hsts override wrong: %+v", out)
	}
	if out.CSPEnabled || !out.CSPReportOnly {
		t.Errorf("csp override wrong: %+v", out)
	}
	if out.XFrameOptions != "SAMEORIGIN" || out.ReferrerPolicy != "no-referrer" {
		t.Errorf("xframe/referrer override wrong: %+v", out)
	}
}
