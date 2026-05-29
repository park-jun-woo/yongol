//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what nonZeroCheckWithCol 단위 테스트 (negated nil-check / native fallback)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestNonZeroCheckWithCol(t *testing.T) {
	cases := []struct {
		name   string
		target string
		col    *ddl.Column
		want   string
	}{
		{"nil col native", "n", nil, "n != 0"},
		{"native col", "n", &ddl.Column{Name: "n", RawType: "BIGINT", NotNull: true}, "n != 0"},
		{"uuid col negated nil-check", "id", &ddl.Column{Name: "id", RawType: "UUID", NotNull: true}, "!pgtypex.IsNilPgUUID(id)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := nonZeroCheckWithCol(tc.target, tc.col); got != tc.want {
				t.Errorf("nonZeroCheckWithCol(%q) = %q, want %q", tc.target, got, tc.want)
			}
		})
	}
}
