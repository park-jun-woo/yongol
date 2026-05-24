//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what isDateColumn — DATE 타입 판정 (매칭/비매칭/대소문자/배열) 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsDateColumn(t *testing.T) {
	cases := []struct {
		name    string
		rawType string
		want    bool
	}{
		{name: "exact_DATE", rawType: "DATE", want: true},
		{name: "lowercase_date", rawType: "date", want: true},
		{name: "date_array", rawType: "DATE[]", want: true},
		{name: "not_date_timestamp", rawType: "TIMESTAMP", want: false},
		{name: "not_date_timestamptz", rawType: "TIMESTAMPTZ", want: false},
		{name: "not_date_text", rawType: "TEXT", want: false},
		{name: "empty", rawType: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			col := ddl.Column{RawType: c.rawType}
			got := isDateColumn(col)
			if got != c.want {
				t.Errorf("isDateColumn(%q) = %v, want %v", c.rawType, got, c.want)
			}
		})
	}
}
