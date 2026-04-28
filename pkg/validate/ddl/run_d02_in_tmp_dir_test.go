//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ddl-structural
//ff:what runD02InTmpDir — D-2 테스트용: 임시 dir 에 sql 쓰고 d02NullableColumn 호출

package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runD02InTmpDir writes sqlContent to <tmp>/db/<fname> and invokes the
// D-2 rule against a Fullstack whose SpecsDir points at <tmp>. Caller
// may pre-populate fs.DDLTables (e.g. to simulate parser-captured
// `-- @nullable` annotations) via the tables argument.
func runD02InTmpDir(t *testing.T, fname, sqlContent string, tables []ddl.Table) []string {
	t.Helper()
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir db: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, fname), []byte(sqlContent), 0o644); err != nil {
		t.Fatalf("write sql: %v", err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp, DDLTables: tables}
	diags := d02NullableColumn(fs)
	msgs := make([]string, len(diags))
	for i, d := range diags {
		msgs[i] = d.Message
	}
	return msgs
}
