//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestCountLeadingSpaces — countLeadingSpaces 가 선행 공백 수를 정확히 세는지 검증

package agent

import "testing"

func TestCountLeadingSpaces(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"abc", 0},
		{"  abc", 2},
		{"    ", 4},
		{"\tabc", 0}, // tab is not a space
		{" \tabc", 1},
	}
	for _, c := range cases {
		if got := countLeadingSpaces(c.in); got != c.want {
			t.Errorf("countLeadingSpaces(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
