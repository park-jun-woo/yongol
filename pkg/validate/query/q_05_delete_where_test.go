//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q05DeleteWhere — Fullstack 단위 DELETE WHERE 검증 (정상/누락/혼합) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ05DeleteWhere(t *testing.T) {
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
		diags := q05DeleteWhere(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("DELETE with WHERE pass", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: DeleteUser :exec\nDELETE FROM users WHERE id = $1;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "DeleteUser", File: f, Line: 1},
			},
		}
		diags := q05DeleteWhere(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("DELETE without WHERE fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: DeleteAll :exec\nDELETE FROM users;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "DeleteAll", File: f, Line: 1},
			},
		}
		diags := q05DeleteWhere(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[Q-05]") {
			t.Errorf("expected Q-05, got %s", diags[0].Message)
		}
	})
}
