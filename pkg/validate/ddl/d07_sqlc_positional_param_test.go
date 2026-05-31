//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD07SqlcPositionalParam(t *testing.T) {
	// scanPositionals reads the actual query file, so write one.
	tmp := t.TempDir()
	qPath := filepath.Join(tmp, "q.sql")
	if err := os.WriteFile(qPath, []byte("-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetUser", File: qPath, Line: 1},
		},
	}
	d := d07SqlcPositionalParam(fs)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[D-7]") {
		t.Errorf("positional param want 1 D-7 diag, got %+v", d)
	}
	// no positional → none
	qPath2 := filepath.Join(tmp, "q2.sql")
	if err := os.WriteFile(qPath2, []byte("-- name: GetUser2 :one\nSELECT * FROM users WHERE id = @id;\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	fs2 := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetUser2", File: qPath2, Line: 1},
		},
	}
	if d := d07SqlcPositionalParam(fs2); len(d) != 0 {
		t.Errorf("named param want 0, got %d", len(d))
	}
}
