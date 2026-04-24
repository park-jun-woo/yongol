//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestDF_06_DBInit_DefersPoolClose — pgx/v5 pool Close + conn Close defer 회귀

package boot

import (
	"strings"
	"testing"
)

// TestDF_06_DBInit_DefersPoolClose — Phase005 pgx/v5 refit. The db-init
// template keeps the DF-06 contract but for the new types: the pgxpool
// Pool gets `defer pool.Close()`. Phase002 (ssac/purify) removed the
// parallel bridged *sql.DB (`conn`) because ssac no longer depends on
// database/sql, so the former `conn.Close()` assertion is dropped. The
// pgxpool.ParseConfig + pgxpool.NewWithConfig error guards remain.
func TestDF_06_DBInit_DefersPoolClose(t *testing.T) {
	block := blockDBInit(nil, "example.com/zenflow")
	lines := strings.Join(block.Lines, "\n")
	if !strings.Contains(lines, "defer pool.Close()") {
		t.Fatalf("db-init must defer pool.Close() (DF-06 on pgx/v5), got:\n%s", lines)
	}
	if strings.Contains(lines, "conn.Close()") {
		t.Fatalf("db-init must NOT retain the ssac database/sql bridge (Phase002 removed it), got:\n%s", lines)
	}
	if !strings.Contains(lines, `pool, err := pgxpool.NewWithConfig(ctx, poolCfg)`) {
		t.Fatalf("db-init must allocate pgxpool via NewWithConfig, got:\n%s", lines)
	}
	if !strings.Contains(lines, "if err != nil {") {
		t.Fatalf("db-init must guard pgxpool errors (DF-01 family), got:\n%s", lines)
	}
}
