//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestFindParamEnd — 경로 파라미터 이름 끝 위치 탐색 (슬래시/문자열끝)
package ssac

import (
	"testing"
)

func TestFindParamEnd(t *testing.T) {
	cases := []struct {
		name  string
		path  string
		start int
		want  int
	}{
		{"StopsAtSlash", "id/sub", 0, 2},
		{"EndOfString", "id", 0, 2},
		{"MidPath", "users/id/edit", 6, 8},
		{"StartAtSlash", "/x", 0, 0},
		{"EmptyString", "", 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := findParamEnd(c.path, c.start); got != c.want {
				t.Errorf("findParamEnd(%q, %d) = %d, want %d", c.path, c.start, got, c.want)
			}
		})
	}
}
