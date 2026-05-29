//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isUUIDColumn 단위 테스트 (UUID 바인딩 판별)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsUUIDColumn(t *testing.T) {
	cases := []struct {
		name string
		col  *ddl.Column
		want bool
	}{
		{"uuid column", &ddl.Column{Name: "id", RawType: "UUID", NotNull: true}, true},
		{"bigint column", &ddl.Column{Name: "id", RawType: "BIGINT", NotNull: true}, false},
		{"text column", &ddl.Column{Name: "name", RawType: "TEXT"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isUUIDColumn(tc.col); got != tc.want {
				t.Errorf("isUUIDColumn(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
