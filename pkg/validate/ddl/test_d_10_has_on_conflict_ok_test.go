//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD10_HasOnConflictOK — ON CONFLICT DO NOTHING 가 있으면 진단 없음
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestD10_HasOnConflictOK asserts an annotated INSERT with the required
// ON CONFLICT DO NOTHING clause produces no diagnostics.
func TestD10_HasOnConflictOK(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	sql := `CREATE TABLE t (id BIGINT);

-- @sentinel
INSERT INTO t (id) VALUES (0) ON CONFLICT DO NOTHING;
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d10SentinelWithoutOnConflict(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
