//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what blockDBInit — tracing 활성 시 otelsql.Open 분기 회귀 방지

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBlockDBInit_OtelWrapsOpen asserts that when tracing is enabled the
// db-init block swaps sql.Open for otelsql.Open so every DB call becomes a
// child span automatically. The non-tracing branch (sql.Open) is covered
// by TestDF_06_DBInit_HasDeferClose; this test guards the tracing branch.
func TestBlockDBInit_OtelWrapsOpen(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: true, Exporter: "stdout"},
				},
			},
		},
	}
	block := blockDBInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")

	if !strings.Contains(body, `otelsql.Open("postgres"`) {
		t.Fatalf("tracing-enabled db-init must use otelsql.Open, got:\n%s", body)
	}
	// The tracing branch must not emit a plain sql.Open — check for the
	// exact non-otel prefix `"\nconn, err := sql.Open("` so we don't false-
	// match otelsql.Open (which is a proper substring).
	if strings.Contains(body, "err := sql.Open(") {
		t.Fatalf("tracing-enabled db-init must NOT fall back to sql.Open, got:\n%s", body)
	}
	if !strings.Contains(imports, "github.com/XSAM/otelsql") {
		t.Fatalf("missing otelsql import, got:\n%s", imports)
	}
	if !strings.Contains(body, "defer conn.Close()") {
		t.Fatalf("tracing branch must still defer conn.Close(), got:\n%s", body)
	}
}

func TestBlockDBInit_NonOtelKeepsSqlOpen(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{},
		},
	}
	block := blockDBInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `sql.Open("postgres"`) {
		t.Fatalf("non-tracing db-init must keep sql.Open, got:\n%s", body)
	}
	if strings.Contains(body, "otelsql") {
		t.Fatalf("non-tracing db-init must NOT import otelsql, got:\n%s", body)
	}
}
