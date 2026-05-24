//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q03NamePascalCase — PascalCase 이름 검증 (정상/소문자/언더스코어/빈 이름) 검증

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ03NamePascalCase(t *testing.T) {
	t.Run("PascalCase names pass", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one"},
				{Name: "ListUsers", Cardinality: "many"},
			},
		}
		diags := q03NamePascalCase(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("lowercase start fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "getUser"},
			},
		}
		diags := q03NamePascalCase(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[Q-03]") {
			t.Errorf("expected Q-03, got %s", diags[0].Message)
		}
	})

	t.Run("underscore fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "Get_User"},
			},
		}
		diags := q03NamePascalCase(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "Get_User") {
			t.Errorf("expected query name in message, got %s", diags[0].Message)
		}
	})

	t.Run("empty name is skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: ""},
			},
		}
		diags := q03NamePascalCase(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("empty queries returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q03NamePascalCase(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
