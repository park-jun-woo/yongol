//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q04CheckManyLimit — :many LIMIT 검사 (pass/fire/non-many/escape/파일 없음) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestQ04CheckManyLimit(t *testing.T) {
	write := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("non-many cardinality returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Cardinality: "one", Name: "GetUser"}
		_, fired := q04CheckManyLimit(q)
		if fired {
			t.Error("expected false for :one cardinality")
		}
	})

	t.Run("exec cardinality returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Cardinality: "exec", Name: "DeleteUser"}
		_, fired := q04CheckManyLimit(q)
		if fired {
			t.Error("expected false for :exec cardinality")
		}
	})

	t.Run("many with LIMIT returns false", func(t *testing.T) {
		f := write(t, "-- name: ListUsers :many\nSELECT * FROM users LIMIT 10;\n")
		q := sqlc.QuerySpec{Cardinality: "many", Name: "ListUsers", File: f, Line: 1}
		_, fired := q04CheckManyLimit(q)
		if fired {
			t.Error("expected false when LIMIT is present")
		}
	})

	t.Run("many with newline LIMIT returns false", func(t *testing.T) {
		f := write(t, "-- name: ListUsers :many\nSELECT * FROM users\nLIMIT 10;\n")
		q := sqlc.QuerySpec{Cardinality: "many", Name: "ListUsers", File: f, Line: 1}
		_, fired := q04CheckManyLimit(q)
		if fired {
			t.Error("expected false when LIMIT on newline is present")
		}
	})

	t.Run("many without LIMIT fires", func(t *testing.T) {
		f := write(t, "-- name: ListUsers :many\nSELECT * FROM users;\n")
		q := sqlc.QuerySpec{Cardinality: "many", Name: "ListUsers", File: f, Line: 1}
		diag, fired := q04CheckManyLimit(q)
		if !fired {
			t.Fatal("expected true for :many without LIMIT")
		}
		if diag.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %v", diag.Level)
		}
		if !strings.Contains(diag.Message, "[Q-04]") {
			t.Errorf("expected Q-04 in message, got %s", diag.Message)
		}
	})

	t.Run("many with @no-pagination escape returns false", func(t *testing.T) {
		f := write(t, "-- name: ListUsers :many\n-- @no-pagination\nSELECT * FROM users;\n")
		q := sqlc.QuerySpec{Cardinality: "many", Name: "ListUsers", File: f, Line: 1}
		_, fired := q04CheckManyLimit(q)
		if fired {
			t.Error("expected false when @no-pagination escape is present")
		}
	})

	t.Run("file not found returns false", func(t *testing.T) {
		q := sqlc.QuerySpec{Cardinality: "many", Name: "ListUsers", File: "/nonexistent/query.sql", Line: 1}
		_, fired := q04CheckManyLimit(q)
		if fired {
			t.Error("expected false when file not found")
		}
	})
}
