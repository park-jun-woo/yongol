//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockDBInit_UsesPgxpool — pgx/v5 refit 후 pgxpool 경로 회귀

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBlockDBInit_UsesPgxpool — Phase005 pgx/v5 refit. The db-init block
// must open a *pgxpool.Pool (not sql.Open) and bridge it to *sql.DB via
// stdlib.OpenDBFromPool so ssac packages keep working. The old sql.Open
// / lib/pq / otelsql branches are removed; otel is to be reinstated later
// via otelpgx.
func TestBlockDBInit_UsesPgxpool(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{},
		},
	}
	block := blockDBInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")

	if !strings.Contains(body, `pgxpool.NewWithConfig`) {
		t.Fatalf("db-init must use pgxpool.NewWithConfig, got:\n%s", body)
	}
	if strings.Contains(body, `sql.Open("postgres"`) {
		t.Fatalf("db-init must NOT use sql.Open (pgx/v5 refit), got:\n%s", body)
	}
	if !strings.Contains(body, `stdlib.OpenDBFromPool(pool)`) {
		t.Fatalf("db-init must bridge pool → *sql.DB via stdlib.OpenDBFromPool, got:\n%s", body)
	}
	if !strings.Contains(body, `queries := db.New(pool)`) {
		t.Fatalf("db-init must initialise sqlc Queries from the pool, got:\n%s", body)
	}
	if !strings.Contains(imports, `github.com/jackc/pgx/v5/pgxpool`) {
		t.Fatalf("missing pgxpool import, got:\n%s", imports)
	}
	if !strings.Contains(imports, `github.com/jackc/pgx/v5/stdlib`) {
		t.Fatalf("missing pgx stdlib import, got:\n%s", imports)
	}
	if strings.Contains(imports, `github.com/lib/pq`) {
		t.Fatalf("lib/pq must NOT be imported in pgx/v5 refit, got:\n%s", imports)
	}
}
