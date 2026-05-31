//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockGinRun — http.Server + SIGTERM graceful shutdown + slog + MaxHeaderBytes
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockGinRun_ManifestHeaderLimit(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{HTTP: &pmanifest.HTTPConfig{HeaderLimit: "2MiB"}},
	}}
	block := blockGinRun(fs)
	if !strings.Contains(strings.Join(block.Lines, "\n"), `"BACKEND_HTTP_HEADER_LIMIT", 2097152`) {
		t.Errorf("manifest header limit not applied, got:\n%v", block.Lines)
	}
}
