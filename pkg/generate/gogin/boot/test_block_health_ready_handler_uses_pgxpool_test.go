//ff:func feature=gen-gogin type=test control=sequence topic=health
//ff:what TestBlockHealth_ReadyHandlerUsesPgxpool — /ready 헬퍼가 pgxpool.Pool 을 사용 (BUG-030)

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockHealth_ReadyHandlerUsesPgxpool — Phase005 pgx/v5 refit made
// *pgxpool.Pool the canonical DB handle. block_db_init emits it as
// `pool`, and the /ready helper must match: signature pgxpool.Pool,
// body pool.Ping, imports pgxpool (no database/sql). Regressions here
// cause the generated backend to fail `go build` with
// `undefined: conn` in main.go. (BUG-030)
func TestBlockHealth_ReadyHandlerUsesPgxpool(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest:  &pmanifest.ProjectConfig{Backend: pmanifest.Backend{}},
		DDLTables: []ddl.Table{{Name: "users"}}, // hasDDL(fs) == true
	}
	block := blockHealth(fs)
	body := strings.Join(block.Lines, "\n")
	funcs := strings.Join(block.Funcs, "\n")
	imports := strings.Join(block.Imports, "\n")

	if !strings.Contains(body, `readyHandlerWithDB(pool)`) {
		t.Fatalf("/ready registration must pass pool (not conn), got:\n%s", body)
	}
	if strings.Contains(body, `readyHandlerWithDB(conn)`) {
		t.Fatalf("/ready registration must NOT reference conn (Phase005 pgx/v5 refit), got:\n%s", body)
	}
	if !strings.Contains(funcs, `readyHandlerWithDB(pool *pgxpool.Pool)`) {
		t.Fatalf("helper signature must take *pgxpool.Pool, got:\n%s", funcs)
	}
	if !strings.Contains(funcs, `pool.Ping(pingCtx)`) {
		t.Fatalf("helper body must call pool.Ping (pgxpool API), got:\n%s", funcs)
	}
	if strings.Contains(funcs, `conn.PingContext`) {
		t.Fatalf("helper must NOT use database/sql PingContext, got:\n%s", funcs)
	}
	if strings.Contains(imports, `"database/sql"`) {
		t.Fatalf("health block must NOT import database/sql (BUG-030), got:\n%s", imports)
	}
	if !strings.Contains(imports, `"github.com/jackc/pgx/v5/pgxpool"`) {
		t.Fatalf("health block must import pgxpool when DDL is present, got:\n%s", imports)
	}
}
