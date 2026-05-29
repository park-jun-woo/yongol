//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockDBInit_UsesPgxpool — pgx/v5 refit 후 pgxpool 경로 회귀 (Phase002 ssac/purify 반영)

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockDBInit_UsesPgxpool — Phase005 pgx/v5 refit established *pgxpool.Pool
// as the primary handle. Phase002 of plans/ssac/purify removed the previous
// stdlib.OpenDBFromPool bridge because ssac no longer depends on database/sql;
// every DB-using ssac adapter is now emitted by yongol codegen
// (pkg/generate/gogin/infra) against the user's sqlc Queries, which already
// wrap the pgxpool.Pool.
//
// This test guards both invariants — pgxpool is the only driver, no
// database/sql bridge survives — so regressions reintroducing sql.Open or
// stdlib.OpenDBFromPool fail fast.
func TestBlockDBInit_UsesPgxpool(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{},
		},
	}
	block := blockDBInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	funcs := strings.Join(block.Funcs, "\n")
	imports := strings.Join(block.Imports, "\n")

	// pgxpool.NewWithConfig lives in the helper func, called from main via initDBPool.
	if !strings.Contains(body, `initDBPool(ctx)`) {
		t.Fatalf("db-init Lines must call initDBPool, got:\n%s", body)
	}
	if !strings.Contains(funcs, `pgxpool.NewWithConfig`) {
		t.Fatalf("db-init helper must use pgxpool.NewWithConfig, got:\n%s", funcs)
	}
	if strings.Contains(funcs, `sql.Open("postgres"`) {
		t.Fatalf("db-init must NOT use sql.Open (pgx/v5 refit), got:\n%s", funcs)
	}
	if strings.Contains(funcs, `stdlib.OpenDBFromPool`) {
		t.Fatalf("db-init must NOT bridge to database/sql (ssac/purify Phase002 removed the bridge), got:\n%s", funcs)
	}
	if !strings.Contains(body, `queries := db.New(pool)`) {
		t.Fatalf("db-init must initialise sqlc Queries from the pool, got:\n%s", body)
	}
	if !strings.Contains(imports, `github.com/jackc/pgx/v5/pgxpool`) {
		t.Fatalf("missing pgxpool import, got:\n%s", imports)
	}
	if strings.Contains(imports, `github.com/jackc/pgx/v5/stdlib`) {
		t.Fatalf("pgx stdlib import must be gone (ssac/purify Phase002), got:\n%s", imports)
	}
	if strings.Contains(imports, `"database/sql"`) {
		t.Fatalf("database/sql import must be gone (ssac/purify Phase002), got:\n%s", imports)
	}
	if strings.Contains(imports, `github.com/lib/pq`) {
		t.Fatalf("lib/pq must NOT be imported in pgx/v5 refit, got:\n%s", imports)
	}
}
