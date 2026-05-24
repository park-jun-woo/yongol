//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q04ManyLimit — Fullstack 단위 :many LIMIT 검증 (정상/누락/혼합) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ04ManyLimit(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("empty queries returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q04ManyLimit(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("many with LIMIT pass", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: ListUsers :many\nSELECT * FROM users LIMIT 10;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "ListUsers", Cardinality: "many", File: f, Line: 1},
			},
		}
		diags := q04ManyLimit(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("many without LIMIT fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: ListAll :many\nSELECT * FROM users;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "ListAll", Cardinality: "many", File: f, Line: 1},
			},
		}
		diags := q04ManyLimit(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[Q-04]") {
			t.Errorf("expected Q-04, got %s", diags[0].Message)
		}
	})

	t.Run("mixed queries only fires for offending", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n-- name: ListAll :many\nSELECT * FROM users;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one", File: f, Line: 1},
				{Name: "ListAll", Cardinality: "many", File: f, Line: 3},
			},
		}
		diags := q04ManyLimit(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})
}
