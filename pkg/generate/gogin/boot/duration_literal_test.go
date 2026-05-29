//ff:func feature=gen-gogin type=test control=sequence
//ff:what durationLiteral — time.Duration을 Go 표현식으로 렌더링 (0 → "0")

package boot

import (
	"testing"
	"time"
)

func TestDurationLiteral(t *testing.T) {
	cases := []struct {
		name string
		in   time.Duration
		want string
	}{
		{"zero", 0, "0"},
		{"one second", time.Second, "time.Duration(1000000000)"},
		{"twelve hours", 12 * time.Hour, "time.Duration(43200000000000)"},
		{"nanosecond", time.Nanosecond, "time.Duration(1)"},
	}
	for _, c := range cases {
		if got := durationLiteral(c.in); got != c.want {
			t.Errorf("%s: durationLiteral(%v) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
