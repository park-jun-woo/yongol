//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what isTimestamptzColumn — TIMESTAMPTZ 판정 (매칭/다중 단어/비매칭) 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsTimestamptzColumn(t *testing.T) {
	cases := []struct {
		name    string
		rawType string
		want    bool
	}{
		{name: "TIMESTAMPTZ", rawType: "TIMESTAMPTZ", want: true},
		{name: "timestamptz_lower", rawType: "timestamptz", want: true},
		{name: "TIMESTAMP_WITH_TIME_ZONE", rawType: "TIMESTAMP WITH TIME ZONE", want: true},
		{name: "timestamptz_array", rawType: "TIMESTAMPTZ[]", want: true},
		{name: "not_timestamp", rawType: "TIMESTAMP", want: false},
		{name: "not_timestamp_without_tz", rawType: "TIMESTAMP WITHOUT TIME ZONE", want: false},
		{name: "not_date", rawType: "DATE", want: false},
		{name: "empty", rawType: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			col := ddl.Column{RawType: c.rawType}
			got := isTimestamptzColumn(col)
			if got != c.want {
				t.Errorf("isTimestamptzColumn(%q) = %v, want %v", c.rawType, got, c.want)
			}
		})
	}
}
