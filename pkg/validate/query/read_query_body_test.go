//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what readQueryBody — SQL 본문 파싱 (정상/빈 파일/파일 없음/복수 쿼리/escape) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestReadQueryBody(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("file not found returns error", func(t *testing.T) {
		q := sqlc.QuerySpec{File: "/nonexistent/query.sql", Line: 1}
		body, err := readQueryBody(q)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if body != nil {
			t.Errorf("expected nil body, got %+v", body)
		}
	})

	t.Run("single query parses body", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n")
		q := sqlc.QuerySpec{File: f, Line: 1}
		body, err := readQueryBody(q)
		if err != nil {
			t.Fatal(err)
		}
		if body == nil {
			t.Fatal("expected body, got nil")
		}
		if body.Header != "-- name: GetUser :one" {
			t.Errorf("expected header, got %q", body.Header)
		}
		if !strings.Contains(body.Text, "SELECT * FROM users") {
			t.Errorf("expected SQL body, got %q", body.Text)
		}
	})

	t.Run("stops at next name annotation", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n-- name: ListUsers :many\nSELECT * FROM users;\n")
		q := sqlc.QuerySpec{File: f, Line: 1}
		body, err := readQueryBody(q)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(body.Text, "ListUsers") {
			t.Errorf("body should not contain next query, got %q", body.Text)
		}
	})

	t.Run("registers escape hatches", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: ListAll :many\nSELECT * FROM users\n-- +no-pagination\n;\n")
		q := sqlc.QuerySpec{File: f, Line: 1}
		body, err := readQueryBody(q)
		if err != nil {
			t.Fatal(err)
		}
		if !body.Escapes["@no-pagination"] {
			t.Error("expected @no-pagination escape to be registered")
		}
		if !body.HasStop {
			t.Error("expected HasStop to be true")
		}
	})

	t.Run("second query in file", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT id FROM users;\n-- name: ListUsers :many\nSELECT * FROM users LIMIT 10;\n")
		q := sqlc.QuerySpec{File: f, Line: 3}
		body, err := readQueryBody(q)
		if err != nil {
			t.Fatal(err)
		}
		if body.Header != "-- name: ListUsers :many" {
			t.Errorf("expected ListUsers header, got %q", body.Header)
		}
		if !strings.Contains(body.Text, "LIMIT 10") {
			t.Errorf("expected LIMIT in body, got %q", body.Text)
		}
	})
}
