//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what pickConvertRHS nullable pgtype 이중 ptrOf 래핑 방지 단위 테스트

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestPickConvertRHS_NullablePgtypeNoDoublePtr(t *testing.T) {
	cases := []struct {
		name        string
		col         *ddl.Column
		isRequired  bool
		wantPrefix  string
		wantNoPtrOf bool
	}{
		{
			name:        "nullable TIMESTAMPTZ + optional → no ptrOf",
			col:         &ddl.Column{RawType: "TIMESTAMPTZ", NotNull: false},
			isRequired:  false,
			wantPrefix:  "pgtypex.FromPgTimestamptzPtr(row.CreatedAt)",
			wantNoPtrOf: true,
		},
		{
			name:        "NOT NULL TIMESTAMPTZ + optional → ptrOf wraps",
			col:         &ddl.Column{RawType: "TIMESTAMPTZ", NotNull: true},
			isRequired:  false,
			wantPrefix:  "ptrOf(pgtypex.FromPgTimestamptz(row.CreatedAt))",
			wantNoPtrOf: false,
		},
		{
			name:        "NOT NULL TIMESTAMPTZ + required → no wrapping",
			col:         &ddl.Column{RawType: "TIMESTAMPTZ", NotNull: true},
			isRequired:  true,
			wantPrefix:  "pgtypex.FromPgTimestamptz(row.CreatedAt)",
			wantNoPtrOf: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := pickConvertRHS("created_at", "CreatedAt", "CreatedAt", tc.isRequired, nil, "", tc.col)
			if got != tc.wantPrefix {
				t.Errorf("pickConvertRHS() = %q, want %q", got, tc.wantPrefix)
			}
		})
	}
}
