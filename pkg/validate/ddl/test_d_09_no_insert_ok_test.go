//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD09_NoInsertOK — INSERT 가 아예 없으면 진단 없음
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestD09_NoInsertOK asserts that a DDL file without any INSERT produces
// no diagnostics.
func TestD09_NoInsertOK(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	sql := `CREATE TABLE t (id BIGINT);
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d09TopLevelInsertWithoutSentinel(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
