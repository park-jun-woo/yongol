//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what blockCsrf 활성/비활성 테스트 (bearer=dormant, cookie=active)

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_BearerMode_Dormant(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Auth:   &pmanifest.Auth{Mode: "bearer"},
			},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("bearer mode should yield inert block, got lines: %+v", block.Lines)
	}
	if block.Active == nil || block.Active(fs) {
		t.Fatalf("bearer mode should report Active()=false")
	}
}

func TestBlockCsrf_NoAuthBlock_Dormant(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("missing auth should yield inert block, got lines: %+v", block.Lines)
	}
	if block.Active == nil || block.Active(fs) {
		t.Fatalf("missing auth should report Active()=false")
	}
}

func TestBlockCsrf_CookieMode_Active(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Auth: &pmanifest.Auth{
					Mode: "cookie",
					Csrf: &pmanifest.CsrfConfig{
						Enabled:     true,
						CookieName:  "XSRF-TOKEN",
						HeaderName:  "X-XSRF-TOKEN",
						ExemptPaths: []string{"/auth/login", "/auth/refresh"},
					},
				},
			},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	if block.Active == nil || !block.Active(fs) {
		t.Fatalf("cookie mode should report Active()=true")
	}
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		"middleware.Csrf(middleware.CsrfConfig{",
		`"XSRF-TOKEN"`,
		`"X-XSRF-TOKEN"`,
		`"/auth/login"`,
		`"/auth/refresh"`,
		"BACKEND_AUTH_CSRF_ENABLED",
		"HybridBearerSkip: false",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("cookie-mode block missing %q, got:\n%s", must, body)
		}
	}
}

func TestBlockCsrf_HybridMode_SetsBearerSkip(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Auth: &pmanifest.Auth{
					Mode: "hybrid",
					Csrf: &pmanifest.CsrfConfig{Enabled: true},
				},
			},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "HybridBearerSkip: true") {
		t.Fatalf("hybrid mode should set HybridBearerSkip:true, got:\n%s", body)
	}
}

func TestHasCsrf_CookieCsrfDisabled_Dormant(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Mode: "cookie",
					Csrf: &pmanifest.CsrfConfig{Enabled: false},
				},
			},
		},
	}
	if hasCsrf(fs) {
		t.Fatalf("csrf.enabled=false must report hasCsrf=false (rejected earlier by SEC-201 at validate)")
	}
}
