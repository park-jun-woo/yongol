//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what lexGuardSymbol 단일 케이스 심볼 토큰/소비 길이/에러 검증 헬퍼

package stml

import "testing"

// assertLexGuardSymbol runs lexGuardSymbol at runes[i] and asserts the token
// kind/text and next index, or the expected error.
func assertLexGuardSymbol(t *testing.T, runes string, i int, wantKind guardTokKind, wantText string, wantNext int, wantErr bool) {
	t.Helper()
	tok, next, err := lexGuardSymbol([]rune(runes), i)
	if wantErr {
		if err == nil {
			t.Fatalf("lexGuardSymbol(%q,%d) expected error", runes, i)
		}
		return
	}
	if err != nil {
		t.Fatalf("lexGuardSymbol(%q,%d) unexpected error: %v", runes, i, err)
	}
	if tok.kind != wantKind || tok.text != wantText {
		t.Errorf("tok = {%d %q}, want {%d %q}", tok.kind, tok.text, wantKind, wantText)
	}
	if next != wantNext {
		t.Errorf("next = %d, want %d", next, wantNext)
	}
}
