//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q06CheckUpdateWhere — UPDATE WHERE 검사 (pass/fire/escape/non-UPDATE/파일 없음) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestQ06CheckUpdateWhere(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("UPDATE with WHERE returns false", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: UpdateUser :exec\nUPDATE users SET name = $1 WHERE id = $2;\n")
		q := sqlc.QuerySpec{Name: "UpdateUser", File: f, Line: 1}
		_, fired := q06CheckUpdateWhere(q)
		if fired {
			t.Error("expected false when WHERE is present")
		}
	})

	t.Run("UPDATE with newline WHERE returns false", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: UpdateUser :exec\nUPDATE users SET name = $1\nWHERE id = $2;\n")
		q := sqlc.QuerySpec{Name: "UpdateUser", File: f, Line: 1}
		_, fired := q06CheckUpdateWhere(q)
		if fired {
			t.Error("expected false when WHERE on newline is present")
		}
	})

	t.Run("UPDATE without WHERE fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: UpdateAll :exec\nUPDATE users SET active = true;\n")
		q := sqlc.QuerySpec{Name: "UpdateAll", File: f, Line: 1}
		diag, fired := q06CheckUpdateWhere(q)
		if !fired {
			t.Fatal("expected true for UPDATE without WHERE")
		}
		if diag.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %v", diag.Level)
		}
		if !strings.Contains(diag.Message, "[Q-06]") {
			t.Errorf("expected Q-06 in message, got %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "UpdateAll") {
			t.Errorf("expected query name in message, got %s", diag.Message)
		}
	})

	t.Run("UPDATE with +allow-truncate escape returns false", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: ResetAll :exec\nUPDATE users SET active = false\n-- +allow-truncate\n;\n")
		q := sqlc.QuerySpec{Name: "ResetAll", File: f, Line: 1}
		_, fired := q06CheckUpdateWhere(q)
		if fired {
			t.Error("expected false when +allow-truncate escape is present")
		}
	})

	t.Run("non-UPDATE query returns false", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n")
		q := sqlc.QuerySpec{Name: "GetUser", File: f, Line: 1}
		_, fired := q06CheckUpdateWhere(q)
		if fired {
			t.Error("expected false for non-UPDATE query")
		}
	})

	t.Run("file not found returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Name: "UpdateUser", File: "/nonexistent/query.sql", Line: 1}
		_, fired := q06CheckUpdateWhere(q)
		if fired {
			t.Error("expected false when file not found")
		}
	})
}
