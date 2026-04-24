//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestDF_06_DBInit_DefersPoolClose — pgx/v5 pool Close + conn Close defer 회귀

package boot

import (
	"strings"
	"testing"
)

// TestDF_06_DBInit_DefersPoolClose — Phase005 pgx/v5 refit. The db-init
// template keeps the DF-06 contract but for the new types: the pgxpool
// Pool gets `defer pool.Close()` and the bridged *sql.DB gets a deferred
// close (wrapped because *sql.DB.Close returns error). Both guard clauses
// on pgxpool.ParseConfig + pgxpool.NewWithConfig errors are asserted so a
// silent os.Exit(1) elision would regress detection here.
func TestDF_06_DBInit_DefersPoolClose(t *testing.T) {
	block := blockDBInit(nil, "example.com/zenflow")
	lines := strings.Join(block.Lines, "\n")
	if !strings.Contains(lines, "defer pool.Close()") {
		t.Fatalf("db-init must defer pool.Close() (DF-06 on pgx/v5), got:\n%s", lines)
	}
	if !strings.Contains(lines, "conn.Close()") {
		t.Fatalf("db-init must close bridged *sql.DB, got:\n%s", lines)
	}
	if !strings.Contains(lines, `pool, err := pgxpool.NewWithConfig(ctx, poolCfg)`) {
		t.Fatalf("db-init must allocate pgxpool via NewWithConfig, got:\n%s", lines)
	}
	if !strings.Contains(lines, "if err != nil {") {
		t.Fatalf("db-init must guard pgxpool errors (DF-01 family), got:\n%s", lines)
	}
}
