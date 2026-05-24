//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what isIntervalColumn — INTERVAL 타입 판정 (매칭/비매칭/대소문자) 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsIntervalColumn(t *testing.T) {
	cases := []struct {
		name    string
		rawType string
		want    bool
	}{
		{name: "INTERVAL", rawType: "INTERVAL", want: true},
		{name: "interval_lower", rawType: "interval", want: true},
		{name: "interval_array", rawType: "INTERVAL[]", want: true},
		{name: "not_interval_text", rawType: "TEXT", want: false},
		{name: "not_interval_timestamp", rawType: "TIMESTAMP", want: false},
		{name: "empty", rawType: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			col := ddl.Column{RawType: c.rawType}
			got := isIntervalColumn(col)
			if got != c.want {
				t.Errorf("isIntervalColumn(%q) = %v, want %v", c.rawType, got, c.want)
			}
		})
	}
}
