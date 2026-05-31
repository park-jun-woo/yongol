//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what wrapJSONBLiteral / looksLikeStringLiteral 단위 테스트 (JSONB 컬럼 string 리터럴 → []byte)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestWrapJSONBLiteral(t *testing.T) {
	g := &methodGen{
		activeMethod: "UserCreate",
		SQLcQueries:  []sqlcparser.QuerySpec{{Name: "UserCreate", Model: "User"}},
		DDLTables: []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"claims": {Name: "claims", RawType: "JSONB", NotNull: true},
					"name":   {Name: "name", RawType: "TEXT", NotNull: true},
				},
			},
		},
	}
	cases := []struct {
		name     string
		inputKey string
		rendered string
		want     string
	}{
		{"jsonb literal wrapped", "Claims", `"{}"`, `[]byte("{}")`},
		{"non-literal untouched", "Claims", "claimsVar", "claimsVar"},
		{"native column untouched", "Name", `"x"`, `"x"`},
		{"empty untouched", "Claims", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.wrapJSONBLiteral(tc.inputKey, tc.rendered); got != tc.want {
				t.Errorf("wrapJSONBLiteral(%q,%q) = %q, want %q", tc.inputKey, tc.rendered, got, tc.want)
			}
		})
	}
}
