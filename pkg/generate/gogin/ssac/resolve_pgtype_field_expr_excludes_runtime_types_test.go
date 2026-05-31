//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what resolvePgtypeFieldExpr 단위 테스트 (dotted 필드의 pgtype 변환식 + 두문자어 컬럼 매칭)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestResolvePgtypeFieldExprExcludesRuntimeTypes(t *testing.T) {
	g := &methodGen{
		VarTypes: map[string]string{"g": "Gadget"},
		DDLTables: []ddl.Table{
			{
				Name: "gadgets",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "UUID", NotNull: true},
				},
			},
		},
	}
	_, imps := g.resolvePgtypeFieldExpr("g.ID")
	for _, imp := range imps {
		if imp == `"github.com/jackc/pgx/v5/pgtype"` {
			t.Errorf("pgtype import must be excluded from handler imports, got %v", imps)
		}
		if imp == `"github.com/oapi-codegen/runtime/types"` {
			t.Errorf("runtime/types import must be excluded from handler imports, got %v", imps)
		}
	}
}
