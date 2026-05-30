//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what trackQuoteState — 라인 스캔 후 단일 인용 상태 / '' 이스케이프

package ddl

import "testing"

func TestTrackQuoteState(t *testing.T) {
	cases := []struct {
		name    string
		line    string
		inStart bool
		want    bool
	}{
		{"balanced", "id = 'x'", false, false},
		{"opens unterminated", "name = 'abc", false, true},
		{"closes from open", "def'", true, false},
		{"escaped pair stays open", "a''b", true, true},
		{"no quotes", "id = 5", false, false},
		{"two quotes toggle off", "''", false, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := trackQuoteState(c.line, c.inStart); got != c.want {
				t.Errorf("trackQuoteState(%q,%v) = %v, want %v", c.line, c.inStart, got, c.want)
			}
		})
	}
}
