//ff:func feature=gen-gogin type=test control=sequence topic=path-convert
//ff:what openAPIPathStep — openAPIPathToGin 한 스텝: 현재 위치에서 '{name}' 또는 단일 바이트 처리

package boot

import (
	"strings"
	"testing"
)

func TestOpenAPIPathStep(t *testing.T) {
	cases := []struct {
		name     string
		p        string
		i        int
		wantNext int
		wantBuf  string
	}{
		{"plain byte", "/ab", 0, 1, "/"},
		{"param segment", "/{id}/x", 1, 5, ":id"},
		{"unterminated brace", "/{id", 1, 2, "{"},
		{"empty braces emit empty param", "/{}", 1, 3, ":"},
	}
	for _, c := range cases {
		var b strings.Builder
		next := openAPIPathStep(c.p, c.i, &b)
		if next != c.wantNext {
			t.Errorf("%s: next = %d, want %d", c.name, next, c.wantNext)
		}
		if b.String() != c.wantBuf {
			t.Errorf("%s: buf = %q, want %q", c.name, b.String(), c.wantBuf)
		}
	}
}
