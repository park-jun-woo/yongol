//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q07SelectStarSensitiveEscape — explicit columns/file not found/escape 검증

package query

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ07SelectStarSensitive_Escape(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	sensitiveTable := ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":       {Name: "id", RawType: "BIGINT"},
			"password": {Name: "password", RawType: "TEXT", Sensitive: true},
		},
	}

	t.Run("SELECT with explicit columns pass", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT id, name FROM users WHERE id = $1;\n")
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{sensitiveTable},
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one", File: f, Line: 1},
			},
		}
		diags := q07SelectStarSensitive(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("file not found skips query", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{sensitiveTable},
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", File: "/nonexistent/query.sql", Line: 1},
			},
		}
		diags := q07SelectStarSensitive(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("escape @allow-sensitive suppresses", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n-- @allow-sensitive\n")
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{sensitiveTable},
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one", File: f, Line: 1},
			},
		}
		diags := q07SelectStarSensitive(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 with escape, got %d: %+v", len(diags), diags)
		}
	})
}
