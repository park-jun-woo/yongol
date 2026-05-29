//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestIsHTTPMethod — YAML 키가 HTTP 메서드인지 판별 검증

package agent

import "testing"

func TestIsHTTPMethod(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"get:", true},
		{"POST:", true},
		{"put: something", true},
		{"delete:", true},
		{"summary:", false},
		{"getters:", false},
		{"", false},
	}
	for _, c := range cases {
		if got := isHTTPMethod(c.in); got != c.want {
			t.Errorf("isHTTPMethod(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
