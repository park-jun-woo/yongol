//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q13PgtypeNumericOverride — NUMERIC override 검증 (nil fs/누락) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ13PgtypeNumericOverride(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := q13PgtypeNumericOverride(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no NUMERIC columns returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SpecsDir: "/tmp/fake",
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT"}}},
			},
		}
		diags := q13PgtypeNumericOverride(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("NUMERIC column without overrides fires Q-13", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte("sql:\n  - gen:\n      go:\n        overrides: []\n"), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir: tmp,
			DDLTables: []ddl.Table{
				{Name: "products", Columns: map[string]ddl.Column{"price": {Name: "price", RawType: "NUMERIC(10,2)"}}},
			},
		}
		diags := q13PgtypeNumericOverride(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "Q-13") {
			t.Errorf("expected Q-13, got %s", diags[0].Message)
		}
	})
}
