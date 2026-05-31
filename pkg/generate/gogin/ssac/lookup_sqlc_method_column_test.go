//ff:func feature=gen-gogin type=test control=sequence
//ff:what lookupSQLCMethodColumn 단위 테스트 (activeMethod → model → DDL 컬럼)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestLookupSQLCMethodColumn(t *testing.T) {
	base := func() *methodGen {
		return &methodGen{
			SQLcQueries: []sqlcparser.QuerySpec{{Name: "WorkflowCreate", Model: "Workflow"}},
			DDLTables: []ddl.Table{
				{
					Name: "workflows",
					Columns: map[string]ddl.Column{
						"title": {Name: "title", RawType: "TEXT"},
					},
				},
			},
		}
	}

	t.Run("no active method → nil", func(t *testing.T) {
		if col := base().lookupSQLCMethodColumn("Title"); col != nil {
			t.Errorf("expected nil, got %+v", col)
		}
	})
	t.Run("active method resolves column", func(t *testing.T) {
		g := base()
		g.activeMethod = "WorkflowCreate"
		col := g.lookupSQLCMethodColumn("Title")
		if col == nil || col.RawType != "TEXT" {
			t.Fatalf("expected TEXT column, got %+v", col)
		}
	})
	t.Run("unknown method → nil", func(t *testing.T) {
		g := base()
		g.activeMethod = "Nope"
		if col := g.lookupSQLCMethodColumn("Title"); col != nil {
			t.Errorf("expected nil, got %+v", col)
		}
	})
}
