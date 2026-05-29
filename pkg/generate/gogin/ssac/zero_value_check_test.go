//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what zeroValueCheckWithCol 단위 테스트 (pgtypex NilCheck 분기 / native fallback)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestZeroValueCheckWithCol(t *testing.T) {
	cases := []struct {
		name   string
		target string
		col    *ddl.Column
		want   string
	}{
		{"nil col native zero", "resourceID", nil, "resourceID == 0"},
		{"native col zero", "n", &ddl.Column{Name: "n", RawType: "BIGINT", NotNull: true}, "n == 0"},
		{"uuid col nil-check", "id", &ddl.Column{Name: "id", RawType: "UUID", NotNull: true}, "pgtypex.IsNilPgUUID(id)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := zeroValueCheckWithCol(tc.target, tc.col); got != tc.want {
				t.Errorf("zeroValueCheckWithCol(%q) = %q, want %q", tc.target, got, tc.want)
			}
		})
	}
}
