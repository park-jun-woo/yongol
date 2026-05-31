//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what zz_zerocov_test — ir.parseBodyLimitOrDefault 0% 커버리지 단위 테스트
package ir

import (
	"testing"
)

func TestParseBodyLimitOrDefault_ZeroCov(t *testing.T) {
	const def = int64(42)
	cases := []struct {
		in   string
		want int64
	}{
		{"", def},
		{"10MiB", 10 * 1024 * 1024},
		{"4KiB", 4 * 1024},
		{"1GiB", 1024 * 1024 * 1024},
		{"5MB", 5 * 1000 * 1000},
		{"3KB", 3 * 1000},
		{"nounit", def}, // default branch (no recognized suffix)
		{"xxMiB", def},  // suffix matches but digits invalid → def
		{"0MiB", def},   // n<=0 → def
		{"MiB", def},    // len not > 3 → default branch
		{"MB", def},     // len not > 2 → default branch
	}
	for _, c := range cases {
		if got := parseBodyLimitOrDefault(c.in, def); got != c.want {
			t.Errorf("parseBodyLimitOrDefault(%q)=%d want %d", c.in, got, c.want)
		}
	}
}
