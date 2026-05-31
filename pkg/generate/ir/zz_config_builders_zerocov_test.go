//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBodyLimitConfig_ZeroCov(t *testing.T) {
	// nil fullstack → defaults.
	if c := buildBodyLimitConfig(nil); c.BodyLimit != 1048576 {
		t.Errorf("nil default body limit = %d", c.BodyLimit)
	}
	// nil manifest.HTTP → defaults.
	fsNoHTTP := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if c := buildBodyLimitConfig(fsNoHTTP); c.MultipartLimit != 33554432 {
		t.Errorf("default multipart = %d", c.MultipartLimit)
	}
	// HTTP with explicit limits → parsed overrides.
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				HTTP: &manifest.HTTPConfig{
					BodyLimit:      "2MiB",
					MultipartLimit: "64MiB",
				},
			},
		},
	}
	c := buildBodyLimitConfig(fs)
	if c.BodyLimit != 2*1048576 {
		t.Errorf("body limit = %d, want %d", c.BodyLimit, 2*1048576)
	}
	if c.MultipartLimit != 64*1048576 {
		t.Errorf("multipart limit = %d, want %d", c.MultipartLimit, 64*1048576)
	}
}

func TestBuildRateLimitConfig_ZeroCov(t *testing.T) {
	if c := buildRateLimitConfig(nil); c != nil {
		t.Errorf("nil fullstack should give nil rate limit")
	}
	// empty rate limit map → nil.
	fsEmpty := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if c := buildRateLimitConfig(fsEmpty); c != nil {
		t.Errorf("empty rate limit should be nil")
	}
	// non-empty → config.
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				RateLimit: manifest.RateLimitConfig{
					"Login": manifest.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		},
	}
	if c := buildRateLimitConfig(fs); c == nil {
		t.Errorf("non-empty rate limit should give config")
	}
}

func TestBuildCORSConfig_ZeroCov(t *testing.T) {
	if c := buildCORSConfig(nil); c != nil {
		t.Errorf("nil → nil cors")
	}
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				CORS: &manifest.CORSConfig{
					AllowOrigins:     []string{"https://x"},
					AllowCredentials: true,
				},
			},
		},
	}
	c := buildCORSConfig(fs)
	if c == nil || !c.AllowCredentials || len(c.AllowOrigins) != 1 {
		t.Errorf("cors config = %+v", c)
	}
}

func TestBuildBearerAuthConfig_ZeroCov(t *testing.T) {
	// not present → nil.
	if c := buildBearerAuthConfig(&prepared.State{}); c != nil {
		t.Errorf("absent auth → nil")
	}
	// present, custom secret + claims.
	ps := &prepared.State{Auth: prepared.Auth{
		Present: true, Mode: "bearer",
		Raw: &manifest.Auth{SecretEnv: "MY_SECRET", Claims: map[string]manifest.ClaimDef{"sub": {}}},
	}}
	c := buildBearerAuthConfig(ps)
	if c == nil || c.SecretEnv != "MY_SECRET" || !c.HasClaims {
		t.Errorf("bearer auth = %+v", c)
	}
	// present, no raw → default secret env.
	ps2 := &prepared.State{Auth: prepared.Auth{Present: true, Mode: "bearer"}}
	if c := buildBearerAuthConfig(ps2); c.SecretEnv != "JWT_SECRET" {
		t.Errorf("default secret env = %q", c.SecretEnv)
	}
}

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

func TestBuildErrorEnvelopeAndValidator_ZeroCov(t *testing.T) {
	if c := buildErrorEnvelopeConfig(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}); c == nil {
		t.Errorf("error envelope should always be non-nil")
	}
	if c := buildRequestValidatorConfig(); !c.Active {
		t.Errorf("request validator should be active")
	}
}

func TestBuildAuthInfraConfig_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	// no raw → all defaults.
	ps := &prepared.State{Auth: prepared.Auth{Present: true, Mode: "cookie"}}
	c := buildAuthInfraConfig(fs, ps)
	if c.SecretEnv != "JWT_SECRET" || c.AccessTokenTTL != "15m" || c.RefreshTokenTTL != "168h" {
		t.Errorf("auth infra defaults = %+v", c)
	}
	// raw with custom values → overrides.
	ps2 := &prepared.State{Auth: prepared.Auth{
		Present: true, Mode: "cookie",
		Raw: &manifest.Auth{SecretEnv: "S", AccessTokenTTL: "5m", RefreshTokenTTL: "24h"},
	}}
	c2 := buildAuthInfraConfig(fs, ps2)
	if c2.SecretEnv != "S" || c2.AccessTokenTTL != "5m" || c2.RefreshTokenTTL != "24h" {
		t.Errorf("auth infra overrides = %+v", c2)
	}
}

func TestBuildActiveBlocks_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	ps := &prepared.State{}
	blocks := buildActiveBlocks(fs, ps)
	if len(blocks) != 25 {
		t.Errorf("expected 25 boot blocks, got %d", len(blocks))
	}
}
