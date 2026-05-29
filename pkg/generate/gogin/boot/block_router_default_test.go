//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what blockRouter — tracing 미활성 시 otelgin 미포함 회귀

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockRouter_Default(t *testing.T) {
	// No tracing → no otelgin import, no middleware registration.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{}},
	}
	block := blockRouter(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")
	if !strings.Contains(body, "gin.Default()") {
		t.Fatalf("router block must create gin engine, got:\n%s", body)
	}
	if strings.Contains(body, "otelgin") {
		t.Fatalf("non-tracing router must NOT register otelgin, got:\n%s", body)
	}
	if strings.Contains(imports, "otelgin") {
		t.Fatalf("non-tracing router must NOT import otelgin, got:\n%s", imports)
	}
}
