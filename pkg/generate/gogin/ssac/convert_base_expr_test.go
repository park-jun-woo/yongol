//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what convertBaseExpr 단위 테스트 (col 유무에 따른 ConvertExpr 확장 / row fallback)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestConvertBaseExpr(t *testing.T) {
	cases := []struct {
		name    string
		dbField string
		col     *ddl.Column
		want    string
	}{
		{"nil col fallback", "Name", nil, "row.Name"},
		{"native col fallback", "Name", &ddl.Column{Name: "name", RawType: "TEXT", NotNull: true}, "row.Name"},
		{"uuid col expanded", "ID", &ddl.Column{Name: "id", RawType: "UUID", NotNull: true}, "pgtypex.FromPgUUID(row.ID)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := convertBaseExpr(tc.dbField, tc.col); got != tc.want {
				t.Errorf("convertBaseExpr(%q) = %q, want %q", tc.dbField, got, tc.want)
			}
		})
	}
}
