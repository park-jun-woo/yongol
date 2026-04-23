//ff:func feature=validate type=test-helper control=sequence topic=ddl-structural
//ff:what runD08InTmpDir — write one DDL file to a tmp specsDir/db/ and run d08SerialTypeBanned
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runD08InTmpDir writes sqlContent to <tmp>/db/<fname> and invokes the
// D-8 rule against a Fullstack whose SpecsDir points at <tmp>.
func runD08InTmpDir(t *testing.T, fname, sqlContent string) []diagnostic.Diagnostic {
	t.Helper()
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir db: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, fname), []byte(sqlContent), 0o644); err != nil {
		t.Fatalf("write sql: %v", err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}
	return d08SerialTypeBanned(fs)
}
