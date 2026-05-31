//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestFormatDuration — 분 미만은 초, 분 이상은 "Nm SSs" 포맷 검증
package agent

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	cases := []struct {
		in   time.Duration
		want string
	}{
		{500 * time.Millisecond, "0.5s"},
		{2 * time.Second, "2.0s"},
		{59 * time.Second, "59.0s"},
		{90 * time.Second, "1m 30s"},
		{2*time.Minute + 5*time.Second, "2m 05s"},
	}
	for _, c := range cases {
		if got := formatDuration(c.in); got != c.want {
			t.Errorf("formatDuration(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
