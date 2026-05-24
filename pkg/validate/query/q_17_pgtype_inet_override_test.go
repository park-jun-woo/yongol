//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q17PgtypeInetOverride — INET override 검증 (nil/누락) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ17PgtypeInetOverride(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := q17PgtypeInetOverride(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("INET column without overrides fires Q-17", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte("sql:\n  - gen:\n      go:\n        overrides: []\n"), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir:  tmp,
			DDLTables: []ddl.Table{{Name: "sessions", Columns: map[string]ddl.Column{"ip_addr": {Name: "ip_addr", RawType: "INET"}}}},
		}
		diags := q17PgtypeInetOverride(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "Q-17") {
			t.Errorf("expected Q-17, got %s", diags[0].Message)
		}
	})
}
