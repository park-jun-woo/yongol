//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isAlreadyPointerApiField 단위 테스트 (ApiField 가 *T 포인터인지)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsAlreadyPointerApiField(t *testing.T) {
	cases := []struct {
		name string
		col  *ddl.Column
		want bool
	}{
		{"nil", nil, false},
		{"nullable varchar → *string", &ddl.Column{Name: "email", RawType: "VARCHAR(255)", NotNull: false}, true},
		{"not null text → string", &ddl.Column{Name: "name", RawType: "TEXT", NotNull: true}, false},
		{"not null bigint → int64", &ddl.Column{Name: "age", RawType: "BIGINT", NotNull: true}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isAlreadyPointerApiField(tc.col); got != tc.want {
				t.Errorf("isAlreadyPointerApiField(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
