//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q05CheckDeleteWhere — DELETE WHERE 검사 (pass/fire/escape/non-DELETE/파일 없음) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestQ05CheckDeleteWhere(t *testing.T) {
	write := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("DELETE with WHERE returns false", func(t *testing.T) {
		f := write(t, "-- name: DeleteUser :exec\nDELETE FROM users WHERE id = $1;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "DeleteUser", File: f, Line: 1}
		_, fired := q05CheckDeleteWhere(q)
		if fired {
			t.Error("expected false when WHERE is present")
		}
	})

	t.Run("DELETE with newline WHERE returns false", func(t *testing.T) {
		f := write(t, "-- name: DeleteUser :exec\nDELETE FROM users\nWHERE id = $1;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "DeleteUser", File: f, Line: 1}
		_, fired := q05CheckDeleteWhere(q)
		if fired {
			t.Error("expected false when WHERE on newline is present")
		}
	})

	t.Run("DELETE without WHERE fires", func(t *testing.T) {
		f := write(t, "-- name: DeleteAll :exec\nDELETE FROM users;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "DeleteAll", File: f, Line: 1}
		diag, fired := q05CheckDeleteWhere(q)
		if !fired {
			t.Fatal("expected true for DELETE without WHERE")
		}
		if diag.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %v", diag.Level)
		}
		if !strings.Contains(diag.Message, "[Q-05]") {
			t.Errorf("expected Q-05 in message, got %s", diag.Message)
		}
	})

	t.Run("DELETE with +allow-truncate escape returns false", func(t *testing.T) {
		f := write(t, "-- name: PurgeAll :exec\nDELETE FROM users\n-- +allow-truncate\n;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "PurgeAll", File: f, Line: 1}
		_, fired := q05CheckDeleteWhere(q)
		if fired {
			t.Error("expected false when +allow-truncate escape is present")
		}
	})

	t.Run("non-DELETE query returns false", func(t *testing.T) {
		f := write(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n")
		q := sqlc.QuerySpec{Cardinality: "one", Name: "GetUser", File: f, Line: 1}
		_, fired := q05CheckDeleteWhere(q)
		if fired {
			t.Error("expected false for non-DELETE query")
		}
	})

	t.Run("file not found returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "DeleteUser", File: "/nonexistent/query.sql", Line: 1}
		_, fired := q05CheckDeleteWhere(q)
		if fired {
			t.Error("expected false when file not found")
		}
	})
}
