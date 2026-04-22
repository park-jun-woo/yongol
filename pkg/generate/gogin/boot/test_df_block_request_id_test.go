//ff:func feature=gen-gogin type=test control=sequence topic=request-id
//ff:what blockRequestID 기본/manifest override 스냅샷

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockRequestID_Defaults(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockRequestID(fs, "example.com/zenflow")
	if block.Name != "request-id" {
		t.Fatalf("unexpected name %q", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		`envBool("BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM", true)`,
		`envString("BACKEND_ERROR_REQUEST_ID_HEADER", "X-Request-Id")`,
		`r.Use(middleware.RequestID(ridTrustUpstream, ridHeader))`,
	} {
		if !strings.Contains(body, must) {
			t.Errorf("request-id block missing %q; body:\n%s", must, body)
		}
	}
}

func TestBlockRequestID_ManifestOverride(t *testing.T) {
	trust := false
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Error: &pmanifest.ErrorConfig{
					RequestID: &pmanifest.RequestIDConfig{
						TrustUpstream: &trust,
						Header:        "X-Trace-Id",
					},
				},
			},
		},
	}
	block := blockRequestID(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `envBool("BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM", false)`) {
		t.Errorf("expected trust_upstream=false default, got:\n%s", body)
	}
	if !strings.Contains(body, `envString("BACKEND_ERROR_REQUEST_ID_HEADER", "X-Trace-Id")`) {
		t.Errorf("expected custom header X-Trace-Id, got:\n%s", body)
	}
}
