//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what castIntegerLiteral 단위 테스트 (정수 리터럴에 Go 타입 캐스트 추가)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
)

func TestCastIntegerLiteral(t *testing.T) {
	g := &methodGen{}
	cases := []struct {
		name     string
		rendered string
		binding  types.GoTypeBinding
		want     string
	}{
		{"int8 literal", "1", types.GoTypeBinding{SqlcGoType: "pgtype.Int8"}, "int64(1)"},
		{"int4 literal", "5", types.GoTypeBinding{SqlcGoType: "pgtype.Int4"}, "int32(5)"},
		{"non-integer rendered untouched", "user.ID", types.GoTypeBinding{SqlcGoType: "pgtype.Int8"}, "user.ID"},
		{"unknown sqlc type no cast", "1", types.GoTypeBinding{SqlcGoType: "pgtype.Text"}, "1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.castIntegerLiteral(tc.rendered, tc.binding); got != tc.want {
				t.Errorf("castIntegerLiteral(%q) = %q, want %q", tc.rendered, got, tc.want)
			}
		})
	}
}
