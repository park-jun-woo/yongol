//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what Run — 전체 Q-* 룰 실행 (nil fs/빈 fs/위반 포함) 검증

package query

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	t.Run("empty fullstack returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("valid queries return no errors for structural rules", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "query.sql")
		os.WriteFile(f, []byte("-- name: GetUser :one\nSELECT id FROM users WHERE id = @user_id;\n"), 0o644)

		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one", File: f, Line: 1, Params: []string{"UserID"}},
			},
		}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("bad cardinality fires Q-02", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "BadQuery", Cardinality: "invalid"},
			},
		}
		diags := Run(fs)
		found := false
		for _, d := range diags {
			if d.Message != "" && len(d.Message) > 5 && d.Message[:5] == "[Q-02" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected Q-02 diagnostic in results, got %+v", diags)
		}
	})
}
