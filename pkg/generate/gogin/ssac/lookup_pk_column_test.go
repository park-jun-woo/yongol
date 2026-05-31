//ff:func feature=gen-gogin type=test control=sequence
//ff:what lookupPKColumn 단위 테스트 (target 변수 → VarTypes 모델 → DDL id 컬럼)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestLookupPKColumn(t *testing.T) {
	g := &methodGen{
		VarTypes: map[string]string{"wf": "Workflow"},
		DDLTables: []ddl.Table{
			{
				Name: "workflows",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "BIGINT"},
				},
			},
		},
	}
	t.Run("resolves via VarTypes", func(t *testing.T) {
		col := g.lookupPKColumn("wf")
		if col == nil || col.RawType != "BIGINT" {
			t.Fatalf("expected BIGINT id column, got %+v", col)
		}
	})
	t.Run("unknown var → nil", func(t *testing.T) {
		if col := g.lookupPKColumn("nope"); col != nil {
			t.Errorf("expected nil, got %+v", col)
		}
	})
}
