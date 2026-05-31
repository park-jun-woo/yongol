//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildSecurityHeadersConfig_ZeroCov(t *testing.T) {
	// default production profile (no manifest overrides).
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	c := buildSecurityHeadersConfig(fs)
	if c == nil || c.Profile != "production" || c.XFrameOptions != "DENY" {
		t.Errorf("default sec headers = %+v", c)
	}
	// full overrides hitting every branch of applySecurityHeadersOverrides.
	fsOv := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				SecurityHeaders: &manifest.SecurityHeadersConfig{
					Profile: "api",
					HSTS: &manifest.HSTSConfig{
						MaxAge: 1000, IncludeSubDomains: true, Preload: true,
					},
					CSP: &manifest.CSPConfig{
						ReportOnly: true,
						Directives: map[string][]string{"default-src": {"'self'"}},
					},
					XFrameOptions:     "SAMEORIGIN",
					ReferrerPolicy:    "no-referrer",
					PermissionsPolicy: map[string][]string{"camera": {}},
				},
			},
		},
	}
	cOv := buildSecurityHeadersConfig(fsOv)
	if cOv.Profile != "api" || cOv.HSTSMaxAge != 1000 || !cOv.HSTSPreload {
		t.Errorf("override hsts = %+v", cOv)
	}
	if cOv.XFrameOptions != "SAMEORIGIN" || cOv.ReferrerPolicy != "no-referrer" {
		t.Errorf("override frame/referrer = %+v", cOv)
	}
	if cOv.CSPEnabled { // api profile disables CSP
		t.Errorf("api profile should disable CSP")
	}
	// dev profile branch → CSPReportOnly.
	fsDev := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				SecurityHeaders: &manifest.SecurityHeadersConfig{Profile: "dev"},
			},
		},
	}
	if cDev := buildSecurityHeadersConfig(fsDev); !cDev.CSPReportOnly {
		t.Errorf("dev profile should set CSPReportOnly")
	}
}
