//ff:func feature=rule type=test control=iteration dimension=1
//ff:what isUpper — ASCII 대문자만 true 반환 (A~Z, 기타 문자·기호 false)

package rule

import "testing"

func TestIsUpper(t *testing.T) {
	cases := []struct {
		in   byte
		want bool
	}{
		{'A', true},
		{'Z', true},
		{'M', true},
		{'a', false},
		{'z', false},
		{'0', false},
		{'9', false},
		{'_', false},
		{' ', false},
		{'@', false}, // just before 'A'
		{'[', false}, // just after 'Z'
	}
	for _, c := range cases {
		if got := isUpper(c.in); got != c.want {
			t.Errorf("isUpper(%q) = %v; want %v", c.in, got, c.want)
		}
	}
}
