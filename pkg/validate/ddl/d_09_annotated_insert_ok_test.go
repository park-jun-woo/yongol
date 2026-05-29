//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD09_AnnotatedInsertOK — @sentinel 가 붙은 INSERT 는 진단 없음
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestD09_AnnotatedInsertOK asserts that an INSERT preceded by
// `-- @sentinel` produces no diagnostics.
func TestD09_AnnotatedInsertOK(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	sql := `CREATE TABLE t (id BIGINT);

-- @sentinel
INSERT INTO t (id) VALUES (0) ON CONFLICT DO NOTHING;
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d09TopLevelInsertWithoutSentinel(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
