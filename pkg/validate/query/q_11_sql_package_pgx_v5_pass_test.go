//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what TestQ11SqlPackagePgxV5_Pass — pgx/v5 는 진단 없음

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ11SqlPackagePgxV5_Pass verifies the required value passes silently.
func TestQ11SqlPackagePgxV5_Pass(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "pgx/v5")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) != 0 {
		t.Fatalf("pgx/v5 must pass Q-11, got %d diagnostics: %+v", len(diags), diags)
	}
}
