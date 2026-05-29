//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD10_UnannotatedIgnored — @sentinel 없는 INSERT 는 D-10 대상 아님 (D-9 담당)
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestD10_UnannotatedIgnored ensures D-10 never fires on an unannotated
// INSERT — that case is owned by D-9.
func TestD10_UnannotatedIgnored(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	// No @sentinel — D-10 should NOT fire (D-9 is responsible for this case)
	sql := `CREATE TABLE t (id BIGINT);

INSERT INTO t (id) VALUES (0);
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d10SentinelWithoutOnConflict(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
