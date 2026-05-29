//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isIdentByte — ASCII letter/digit/underscore 판정 (import usage 단어경계)

package boot

import "testing"

func TestIsIdentByte(t *testing.T) {
	cases := []struct {
		name string
		b    byte
		want bool
	}{
		{"lower a", 'a', true},
		{"lower z", 'z', true},
		{"upper A", 'A', true},
		{"upper Z", 'Z', true},
		{"digit 0", '0', true},
		{"digit 9", '9', true},
		{"underscore", '_', true},
		{"dot", '.', false},
		{"space", ' ', false},
		{"slash", '/', false},
		{"quote", '"', false},
		{"hyphen", '-', false},
	}
	for _, c := range cases {
		if got := isIdentByte(c.b); got != c.want {
			t.Errorf("%s: isIdentByte(%q) = %v, want %v", c.name, c.b, got, c.want)
		}
	}
}
