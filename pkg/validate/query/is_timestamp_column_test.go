//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what isTimestampColumn — TIMESTAMP (no TZ) 판정 (매칭/비매칭/다중 단어 정규화) 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsTimestampColumn(t *testing.T) {
	cases := []struct {
		name    string
		rawType string
		want    bool
	}{
		{name: "TIMESTAMP", rawType: "TIMESTAMP", want: true},
		{name: "timestamp_lower", rawType: "timestamp", want: true},
		{name: "TIMESTAMP_WITHOUT_TIME_ZONE", rawType: "TIMESTAMP WITHOUT TIME ZONE", want: true},
		{name: "timestamp_array", rawType: "TIMESTAMP[]", want: true},
		{name: "not_timestamptz", rawType: "TIMESTAMPTZ", want: false},
		{name: "not_timestamp_with_tz", rawType: "TIMESTAMP WITH TIME ZONE", want: false},
		{name: "not_date", rawType: "DATE", want: false},
		{name: "empty", rawType: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			col := ddl.Column{RawType: c.rawType}
			got := isTimestampColumn(col)
			if got != c.want {
				t.Errorf("isTimestampColumn(%q) = %v, want %v", c.rawType, got, c.want)
			}
		})
	}
}
