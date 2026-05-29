//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestSpliceLines — 라인 범위 교체 및 경계/역순 인자 처리 검증

package agent

import "testing"

func TestSpliceLines(t *testing.T) {
	cases := []struct {
		name    string
		content string
		start   int
		end     int
		repl    string
		want    string
	}{
		{name: "replace middle", content: "a\nb\nc\nd", start: 1, end: 3, repl: "X", want: "a\nX\nd"},
		{name: "replace with empty drops range", content: "a\nb\nc", start: 1, end: 2, repl: "", want: "a\nc"},
		{name: "negative start clamps to 0", content: "a\nb\nc", start: -5, end: 1, repl: "X", want: "X\nb\nc"},
		{name: "end past length clamps", content: "a\nb", start: 1, end: 99, repl: "X", want: "a\nX\n"},
		{name: "start greater than end returns original", content: "a\nb\nc", start: 3, end: 1, repl: "X", want: "a\nb\nc"},
		{name: "trailing newline in replacement trimmed", content: "a\nb", start: 0, end: 1, repl: "X\n\n", want: "X\nb"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := spliceLines(c.content, c.start, c.end, c.repl); got != c.want {
				t.Errorf("spliceLines = %q, want %q", got, c.want)
			}
		})
	}
}
