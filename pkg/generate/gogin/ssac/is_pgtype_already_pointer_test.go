//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isPgtypeAlreadyPointer 단위 테스트 (dotted 필드의 pgtype 변환이 이미 포인터인지)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsPgtypeAlreadyPointer(t *testing.T) {
	g := &methodGen{
		VarTypes: map[string]string{
			"user": "User",
			"list": "[]User",
		},
		DDLTables: []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"email": {Name: "email", RawType: "VARCHAR(255)", NotNull: false}, // → *string
					"name":  {Name: "name", RawType: "TEXT", NotNull: true},           // → string
				},
			},
		},
	}
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"no dot", "user", false},
		{"unknown var", "foo.Email", false},
		{"nullable column pointer", "user.Email", true},
		{"not null column non-pointer", "user.Name", false},
		{"slice type stripped, pointer col", "list.Email", true},
		{"missing column", "user.Missing", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.isPgtypeAlreadyPointer(tc.in); got != tc.want {
				t.Errorf("isPgtypeAlreadyPointer(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
