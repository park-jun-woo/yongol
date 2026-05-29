//ff:func feature=gen-gogin type=test control=sequence topic=request-id
//ff:what blockRequestID — middleware.RequestID(...) 최상위 등록 (Phase004)

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRequestID_DefaultsAndImports(t *testing.T) {
	block := blockRequestID(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		`envBool("BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM", true)`,
		`envString("BACKEND_ERROR_REQUEST_ID_HEADER", "X-Request-Id")`,
		"r.Use(middleware.RequestID(ridTrustUpstream, ridHeader))",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("blockRequestID missing %q, got:\n%s", must, body)
		}
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `"example.com/zenflow/internal/middleware"`) {
		t.Errorf("must import middleware, got:\n%v", block.Imports)
	}
}
