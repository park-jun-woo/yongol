//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what lexGuardIdent 단일 케이스 토큰/소비 길이 검증 헬퍼

package stml

import "testing"

// assertLexGuardIdent runs lexGuardIdent on runes[start:] and asserts the token
// text and next index match expectations.
func assertLexGuardIdent(t *testing.T, runes string, start int, wantText string, wantNext int) {
	t.Helper()
	tok, next := lexGuardIdent([]rune(runes), start)
	if tok.kind != tokIdent {
		t.Errorf("kind = %d, want tokIdent", tok.kind)
	}
	if tok.text != wantText {
		t.Errorf("text = %q, want %q", tok.text, wantText)
	}
	if next != wantNext {
		t.Errorf("next = %d, want %d", next, wantNext)
	}
}
