//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what isInetColumn — INET/CIDR 타입 판정 (매칭/비매칭/대소문자) 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsInetColumn(t *testing.T) {
	cases := []struct {
		name    string
		rawType string
		want    bool
	}{
		{name: "INET", rawType: "INET", want: true},
		{name: "inet_lower", rawType: "inet", want: true},
		{name: "CIDR", rawType: "CIDR", want: true},
		{name: "cidr_lower", rawType: "cidr", want: true},
		{name: "inet_array", rawType: "INET[]", want: true},
		{name: "cidr_array", rawType: "CIDR[]", want: true},
		{name: "not_inet_text", rawType: "TEXT", want: false},
		{name: "not_inet_uuid", rawType: "UUID", want: false},
		{name: "empty", rawType: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			col := ddl.Column{RawType: c.rawType}
			got := isInetColumn(col)
			if got != c.want {
				t.Errorf("isInetColumn(%q) = %v, want %v", c.rawType, got, c.want)
			}
		})
	}
}
