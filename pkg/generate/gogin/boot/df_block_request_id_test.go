//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=request-id
//ff:what TestBlockRequestID_Defaults — 기본 request-id 미들웨어 스냅샷

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
