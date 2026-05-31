//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockDBInit_TracingBranchRemoved — Phase005 pgx/v5 refit 로 otelsql 분기 제거

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockDBInit_TracingBranchRemoved — Phase005 pgx/v5 refit deleted the
// otelsql-wrapped code path. Tracing over pgx is a follow-up (obs01 /
// sqlc02) via otelpgx. This test makes sure the manifest "tracing=enabled"
// shape no longer produces an otelsql import — otherwise we'd ship a
// vestigial reference that breaks the Phase006 grep sweep.
func TestBlockDBInit_TracingBranchRemoved(t *testing.T) {
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
	funcs := strings.Join(block.Funcs, "\n")
	imports := strings.Join(block.Imports, "\n")

	if strings.Contains(imports, "github.com/XSAM/otelsql") {
		t.Fatalf("otelsql import must be gone after Phase005, got:\n%s", imports)
	}
	if strings.Contains(funcs, "otelsql.") {
		t.Fatalf("otelsql reference must be gone after Phase005, got:\n%s", funcs)
	}
	if !strings.Contains(funcs, "pgxpool.NewWithConfig") {
		t.Fatalf("tracing-enabled db-init must still take the pgxpool path, got:\n%s", funcs)
	}
}
