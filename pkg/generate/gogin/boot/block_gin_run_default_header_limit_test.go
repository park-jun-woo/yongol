//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockGinRun — http.Server + SIGTERM graceful shutdown + slog + MaxHeaderBytes
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockGinRun_DefaultHeaderLimit(t *testing.T) {
	block := blockGinRun(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}})
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `envInt64("BACKEND_HTTP_HEADER_LIMIT", 1048576)`) {
		t.Errorf("default header limit (1MiB) wrong, got:\n%s", body)
	}
	if !strings.Contains(body, `envInt("BACKEND_HTTP_PORT", 8080)`) {
		t.Errorf("default port (8080) wrong, got:\n%s", body)
	}
	if !strings.Contains(body, "runServerWithGracefulShutdown(r, cancelBootstrap, port, int(headerLimit))") {
		t.Errorf("must delegate to lifecycle helper with port, got:\n%s", body)
	}
	if len(block.Funcs) != 1 {
		t.Errorf("must emit lifecycle helper func, got %d", len(block.Funcs))
	}
}
