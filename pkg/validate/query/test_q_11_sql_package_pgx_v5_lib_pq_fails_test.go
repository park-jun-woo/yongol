//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what TestQ11SqlPackagePgxV5_LibPqFails — lib/pq 는 Q-11 에러

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ11SqlPackagePgxV5_LibPqFails asserts that lib/pq fires Q-11.
func TestQ11SqlPackagePgxV5_LibPqFails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "lib/pq")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("lib/pq must fire Q-11, got no diagnostics")
	}
}
