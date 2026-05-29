//ff:func feature=migration type=test control=selection dimension=2
//ff:what TestIsIdentOrSpaceRune — 식별자 문자/공백/_ 은 true, 그 외 false
package migration

import "testing"

func TestIsIdentOrSpaceRune(t *testing.T) {
	for _, r := range []rune{'a', 'Z', '5', '_', ' '} {
		if !isIdentOrSpaceRune(r) {
			t.Errorf("isIdentOrSpaceRune(%q) = false, want true", r)
		}
	}
	for _, r := range []rune{'-', '+', '(', '.'} {
		if isIdentOrSpaceRune(r) {
			t.Errorf("isIdentOrSpaceRune(%q) = true, want false", r)
		}
	}
}
