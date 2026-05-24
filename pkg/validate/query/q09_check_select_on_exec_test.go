//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q09CheckSelectOnExec — :exec SELECT/RETURNING 검출 (fire/pass/non-exec/파일 없음) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestQ09CheckSelectOnExec(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("non-exec cardinality returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Cardinality: "one", Name: "GetUser"}
		_, fired := q09CheckSelectOnExec(q)
		if fired {
			t.Error("expected false for :one cardinality")
		}
	})

	t.Run("exec DELETE without SELECT/RETURNING returns false", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: DeleteUser :exec\nDELETE FROM users WHERE id = $1;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "DeleteUser", File: f, Line: 1}
		_, fired := q09CheckSelectOnExec(q)
		if fired {
			t.Error("expected false for DELETE without RETURNING")
		}
	})

	t.Run("exec with top-level SELECT fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: BadExec :exec\nSELECT * FROM users;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "BadExec", File: f, Line: 1}
		diag, fired := q09CheckSelectOnExec(q)
		if !fired {
			t.Fatal("expected true for :exec with SELECT")
		}
		if diag.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %v", diag.Level)
		}
		if !strings.Contains(diag.Message, "[Q-09]") {
			t.Errorf("expected Q-09 in message, got %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "BadExec") {
			t.Errorf("expected query name in message, got %s", diag.Message)
		}
	})

	t.Run("exec with RETURNING fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: InsertUser :exec\nINSERT INTO users (name) VALUES ($1) RETURNING id;\n")
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "InsertUser", File: f, Line: 1}
		diag, fired := q09CheckSelectOnExec(q)
		if !fired {
			t.Fatal("expected true for :exec with RETURNING")
		}
		if !strings.Contains(diag.Message, "[Q-09]") {
			t.Errorf("expected Q-09 in message, got %s", diag.Message)
		}
	})

	t.Run("file not found returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "BadExec", File: "/nonexistent/query.sql", Line: 1}
		_, fired := q09CheckSelectOnExec(q)
		if fired {
			t.Error("expected false when file not found")
		}
	})
}
