//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q01NameRequired — nil fs/빈 쿼리/정상 파일/누락 파일 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ01NameRequired(t *testing.T) {
	writeSQLFile := func(t *testing.T, name, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := q01NameRequired(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no queries returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q01NameRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("properly annotated file returns nil", func(t *testing.T) {
		f := writeSQLFile(t, "query.sql", "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", File: f, Line: 1},
			},
		}
		diags := q01NameRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("SQL without name fires diagnostic", func(t *testing.T) {
		// A file with only bare SQL and no -- name: annotation at all
		f := writeSQLFile(t, "bare.sql", "SELECT 1;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "Dummy", File: f, Line: 1},
			},
		}
		diags := q01NameRequired(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[Q-01]") {
			t.Errorf("expected Q-01 in message, got %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "bare.sql") {
			t.Errorf("expected file name in message, got %s", diags[0].Message)
		}
	})

	t.Run("duplicate file refs deduplicated", func(t *testing.T) {
		f := writeSQLFile(t, "query.sql", "SELECT 1;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "A", File: f, Line: 1},
				{Name: "B", File: f, Line: 1},
			},
		}
		diags := q01NameRequired(fs)
		// Only one diagnostic per file even with multiple queries referencing it
		if len(diags) != 1 {
			t.Fatalf("expected 1 (deduplicated), got %d: %+v", len(diags), diags)
		}
	})
}
