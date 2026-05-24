//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q08UnusedParam — 미사용 파라미터 검출 (정상/누락/파라미터 없음/파일 없음) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ08UnusedParam(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("no params skips", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetAll", Params: nil},
			},
		}
		diags := q08UnusedParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("all params referenced pass", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = @user_id;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Params: []string{"UserID"}, File: f, Line: 1},
			},
		}
		diags := q08UnusedParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("sqlc.arg form referenced pass", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = sqlc.arg(user_id);\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Params: []string{"UserID"}, File: f, Line: 1},
			},
		}
		diags := q08UnusedParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("unused param fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = @user_id;\n")
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Params: []string{"UserID", "OrgID"}, File: f, Line: 1},
			},
		}
		diags := q08UnusedParam(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[Q-08]") {
			t.Errorf("expected Q-08, got %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "OrgID") {
			t.Errorf("expected param name OrgID, got %s", diags[0].Message)
		}
	})

	t.Run("file not found skips", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Params: []string{"UserID"}, File: "/nonexistent/query.sql", Line: 1},
			},
		}
		diags := q08UnusedParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("empty queries returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q08UnusedParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
