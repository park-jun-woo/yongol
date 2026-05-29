//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestIsSpaceByte — 공백류(' '/'\t'/'\n'/'\r')만 true
package migration

import "testing"

func TestIsSpaceByte(t *testing.T) {
	for _, c := range []byte{' ', '\t', '\n', '\r'} {
		if !isSpaceByte(c) {
			t.Errorf("isSpaceByte(%q) = false, want true", c)
		}
	}
	for _, c := range []byte{'a', '0', '_', ';'} {
		if isSpaceByte(c) {
			t.Errorf("isSpaceByte(%q) = true, want false", c)
		}
	}
}
