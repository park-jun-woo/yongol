//ff:func feature=gen-gogin type=test control=sequence
//ff:what lookupResourcePKColumn 단위 테스트 (리소스명 → DDL id 컬럼)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestLookupResourcePKColumn(t *testing.T) {
	g := &methodGen{
		DDLTables: []ddl.Table{
			{
				Name: "workflows",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "UUID"},
				},
			},
		},
	}
	t.Run("resolves id column", func(t *testing.T) {
		col := g.lookupResourcePKColumn("workflow")
		if col == nil {
			t.Fatal("expected id column, got nil")
		}
		if col.RawType != "UUID" {
			t.Errorf("RawType = %q, want UUID", col.RawType)
		}
	})
	t.Run("unknown resource → nil", func(t *testing.T) {
		if col := g.lookupResourcePKColumn("ghost"); col != nil {
			t.Errorf("expected nil, got %+v", col)
		}
	})
}
