//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=query-structural
//ff:what loadSpecs — testdata/*.sql 을 실제 sqlc 파서로 파싱해 QuerySpec 반환

package query

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// loadSpecs parses a testdata/*.sql file via the real sqlc parser and returns
// the QuerySpec slice with absolute file paths (required by readQueryBody).
func loadSpecs(t *testing.T, name string) []sqlc.QuerySpec {
	t.Helper()
	abs, err := filepath.Abs(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("abs: %v", err)
	}
	specs, diags := sqlc.ParseFile(abs)
	for _, d := range diags {
		if d.Level == "ERROR" {
			t.Fatalf("parse diag: %+v", d)
		}
	}
	return specs
}
