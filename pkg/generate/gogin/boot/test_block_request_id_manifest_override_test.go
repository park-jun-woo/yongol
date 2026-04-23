//ff:func feature=gen-gogin type=test control=sequence topic=request-id
//ff:what TestBlockRequestID_ManifestOverride — trust_upstream/header override 전달 확인

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
