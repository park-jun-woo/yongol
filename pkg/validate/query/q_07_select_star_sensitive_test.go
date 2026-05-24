//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q07SelectStarSensitive — SELECT * @sensitive 컬럼 검출 (no table/fire/pass) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ07SelectStarSensitive(t *testing.T) {
	writeSQLFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	sensitiveTable := ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":       {Name: "id", RawType: "BIGINT"},
			"password": {Name: "password", RawType: "TEXT", Sensitive: true},
		},
	}

	t.Run("no DDL tables returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q07SelectStarSensitive(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("SELECT * on sensitive table fires", func(t *testing.T) {
		f := writeSQLFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n")
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{sensitiveTable},
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one", File: f, Line: 1},
			},
		}
		diags := q07SelectStarSensitive(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[Q-07]") {
			t.Errorf("expected Q-07, got %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "password") {
			t.Errorf("expected sensitive column name, got %s", diags[0].Message)
		}
	})

	t.Run("SELECT * on non-sensitive table pass", func(t *testing.T) {
		noSensitive := ddl.Table{
			Name: "items",
			Columns: map[string]ddl.Column{
				"id":   {Name: "id", RawType: "BIGINT"},
				"name": {Name: "name", RawType: "TEXT"},
			},
		}
		f := writeSQLFile(t, "-- name: ListItems :many\nSELECT * FROM items LIMIT 10;\n")
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{noSensitive},
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "ListItems", Cardinality: "many", File: f, Line: 1},
			},
		}
		diags := q07SelectStarSensitive(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
