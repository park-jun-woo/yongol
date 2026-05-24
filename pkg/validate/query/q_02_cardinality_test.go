//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q02Cardinality — 유효/무효/빈 cardinality 검증

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ02Cardinality(t *testing.T) {
	t.Run("valid cardinalities return nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser", Cardinality: "one"},
				{Name: "ListUsers", Cardinality: "many"},
				{Name: "DeleteUser", Cardinality: "exec"},
				{Name: "CountUsers", Cardinality: "execrows"},
				{Name: "InsertUser", Cardinality: "execresult"},
				{Name: "LastInsert", Cardinality: "execlastid"},
			},
		}
		diags := q02Cardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("empty cardinality fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "BadQuery", Cardinality: ""},
			},
		}
		diags := q02Cardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[Q-02]") {
			t.Errorf("expected Q-02, got %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "BadQuery") {
			t.Errorf("expected query name, got %s", diags[0].Message)
		}
	})

	t.Run("invalid cardinality fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "WrongQuery", Cardinality: "batch"},
			},
		}
		diags := q02Cardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "batch") {
			t.Errorf("expected cardinality in message, got %s", diags[0].Message)
		}
	})

	t.Run("empty queries returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q02Cardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
