//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what lexGuardString 단일 케이스 문자열 토큰/소비 길이/에러 검증 헬퍼

package stml

import "testing"

// assertLexGuardString runs lexGuardString at runes[i] and asserts the token
// text and next index, or the expected error.
func assertLexGuardString(t *testing.T, runes string, i int, wantText string, wantNext int, wantErr bool) {
	t.Helper()
	tok, next, err := lexGuardString([]rune(runes), i)
	if wantErr {
		if err == nil {
			t.Fatalf("lexGuardString(%q,%d) expected error", runes, i)
		}
		return
	}
	if err != nil {
		t.Fatalf("lexGuardString(%q,%d) unexpected error: %v", runes, i, err)
	}
	if tok.kind != tokString {
		t.Errorf("kind = %d, want tokString", tok.kind)
	}
	if tok.text != wantText {
		t.Errorf("text = %q, want %q", tok.text, wantText)
	}
	if next != wantNext {
		t.Errorf("next = %d, want %d", next, wantNext)
	}
}
