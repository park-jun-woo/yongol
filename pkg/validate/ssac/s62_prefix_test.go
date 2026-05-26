//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what s62prefix — 도트 이전 접두사 추출 검증 (정상/도트 없음/도트가 처음/다중 도트)

package ssac

import "testing"

func TestS62Prefix(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "with dot", in: "course.ID", want: "course"},
		{name: "no dot", in: "course", want: "course"},
		{name: "dot at start", in: ".field", want: ".field"},
		{name: "multiple dots", in: "a.b.c", want: "a"},
		{name: "empty", in: "", want: ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := s62prefix(c.in)
			if got != c.want {
				t.Errorf("s62prefix(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
