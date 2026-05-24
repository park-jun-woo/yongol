//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q15PgtypeTimestampOverride — TIMESTAMP override 검증 (nil/누락) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ15PgtypeTimestampOverride(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := q15PgtypeTimestampOverride(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("TIMESTAMP column without overrides fires Q-15", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte("sql:\n  - gen:\n      go:\n        overrides: []\n"), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir:  tmp,
			DDLTables: []ddl.Table{{Name: "logs", Columns: map[string]ddl.Column{"logged_at": {Name: "logged_at", RawType: "TIMESTAMP"}}}},
		}
		diags := q15PgtypeTimestampOverride(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "Q-15") {
			t.Errorf("expected Q-15, got %s", diags[0].Message)
		}
	})
}
