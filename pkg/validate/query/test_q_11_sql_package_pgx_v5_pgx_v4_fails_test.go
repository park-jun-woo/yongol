//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what TestQ11SqlPackagePgxV5_PgxV4Fails — pgx/v4 는 Q-11 에러

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ11SqlPackagePgxV5_PgxV4Fails asserts that pgx/v4 fires Q-11.
func TestQ11SqlPackagePgxV5_PgxV4Fails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "pgx/v4")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("pgx/v4 must fire Q-11, got no diagnostics")
	}
}
