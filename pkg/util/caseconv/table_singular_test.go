//ff:func feature=util type=test control=iteration dimension=1 topic=string-convert
//ff:what TableSingular 단위 테스트 — ies/sses/xes/s/기본·단수입력 멱등

package caseconv

import "testing"

func TestTableSingular(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"ies suffix", "companies", "company"},
		{"sses suffix", "addresses", "address"},
		{"xes suffix", "boxes", "box"},
		{"plain s suffix", "users", "user"},
		{"plural snake", "bid_requests", "bid_request"},
		{"ss not stripped", "address", "address"},
		{"already singular", "app_config", "app_config"},
		{"singular idempotent", "user", "user"},
		{"y singular idempotent", "company", "company"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := TableSingular(c.in); got != c.want {
				t.Errorf("TableSingular(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
