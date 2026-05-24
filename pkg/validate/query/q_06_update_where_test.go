//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q06UpdateWhere — Fullstack 단위 UPDATE WHERE 검증 (정상/누락/혼합) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ06UpdateWhere(t *testing.T) {
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
		diags := q06UpdateWhere(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("UPDATE with WHERE pass", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: UpdateUser :exec\nUPDATE users SET name = $1 WHERE id = $2;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "UpdateUser", File: f, Line: 1},
			},
		}
		diags := q06UpdateWhere(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("UPDATE without WHERE fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: UpdateAll :exec\nUPDATE users SET active = true;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "UpdateAll", File: f, Line: 1},
			},
		}
		diags := q06UpdateWhere(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[Q-06]") {
			t.Errorf("expected Q-06, got %s", diags[0].Message)
		}
	})
}
