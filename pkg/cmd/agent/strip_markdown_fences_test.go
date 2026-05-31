//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestStripMarkdownFences — 감싸는 markdown 코드 펜스 제거 검증
package agent

import (
	"testing"
)

func TestStripMarkdownFences(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"no fence", "plain", "plain"},
		{"yaml fence", "```yaml\nkey: v\n```", "key: v"},
		{"bare fence", "```\nbody\n```", "body"},
		{"ws trimmed", "\n```\nx\n```\n", "x"},
		{"fence no newline", "```yaml", "```yaml"},
		{"opening fence no closing", "```\nbody", "body"},
	}
	for _, c := range cases {
		if got := stripMarkdownFences(c.in); got != c.want {
			t.Errorf("%s: stripMarkdownFences(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
