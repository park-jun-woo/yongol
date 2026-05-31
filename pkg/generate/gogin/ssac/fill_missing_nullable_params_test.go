//ff:func feature=gen-gogin type=test control=sequence
//ff:what fillMissingNullableParams 단위 테스트 (Inputs에 없는 nullable pgtype param을 zero 값으로 채움)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestFillMissingNullableParams(t *testing.T) {
	g := &methodGen{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "NoteCreate", Model: "Note", Params: []string{"title", "body"}},
		},
		DDLTables: []ddl.Table{
			{
				Name: "notes",
				Columns: map[string]ddl.Column{
					"title": {Name: "title", RawType: "TEXT", NotNull: true}, // native, required
					"body":  {Name: "body", RawType: "TEXT", NotNull: false}, // nullable pgtype
				},
			},
		},
	}

	t.Run("unknown method → nil", func(t *testing.T) {
		fields, imps := g.fillMissingNullableParams("Ghost", map[string]string{})
		if fields != nil || imps != nil {
			t.Errorf("expected nil/nil, got %v / %v", fields, imps)
		}
	})

	t.Run("fills nullable param missing from inputs", func(t *testing.T) {
		// title provided; body omitted and nullable → zero-value literal.
		fields, imps := g.fillMissingNullableParams("NoteCreate", map[string]string{"title": "request.title"})
		if len(fields) != 1 {
			t.Fatalf("expected 1 field, got %v", fields)
		}
		if fields[0] != "body: pgtype.Text{}" {
			t.Errorf("field = %q, want body: pgtype.Text{}", fields[0])
		}
		if len(imps) != 1 || imps[0] != `"github.com/jackc/pgx/v5/pgtype"` {
			t.Errorf("imports = %v, want pgtype", imps)
		}
	})

	t.Run("all params provided → no fills", func(t *testing.T) {
		fields, imps := g.fillMissingNullableParams("NoteCreate", map[string]string{
			"title": "request.title",
			"body":  "request.body",
		})
		if fields != nil || imps != nil {
			t.Errorf("expected no fills, got %v / %v", fields, imps)
		}
	})
}
