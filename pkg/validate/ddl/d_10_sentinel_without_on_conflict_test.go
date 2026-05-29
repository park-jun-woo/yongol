//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD10_MissingOnConflictFlagged — `@sentinel` INSERT 에 ON CONFLICT DO NOTHING 없으면 D-10 ERROR
package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestD10_MissingOnConflictFlagged asserts that an `-- @sentinel` INSERT
// missing the ON CONFLICT DO NOTHING clause raises one D-10 diagnostic.
func TestD10_MissingOnConflictFlagged(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	sql := `CREATE TABLE t (id BIGINT);

-- @sentinel
INSERT INTO t (id) VALUES (0);
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d10SentinelWithoutOnConflict(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[D-10]") {
		t.Errorf("expected [D-10] prefix: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Advice, "ON CONFLICT DO NOTHING") {
		t.Errorf("advice should mention ON CONFLICT DO NOTHING: %q", diags[0].Advice)
	}
}
