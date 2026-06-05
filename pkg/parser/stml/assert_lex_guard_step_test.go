//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what lexGuardStep 단일 케이스 토큰/emit/next/에러 검증 헬퍼

package stml

import "testing"

// assertLexGuardStep runs lexGuardStep at runes[i] and asserts emit/next/token
// or the expected error.
func assertLexGuardStep(t *testing.T, runes string, i int, wantKind guardTokKind, wantText string, wantEmit bool, wantNext int, wantErr bool) {
	t.Helper()
	tok, emit, next, err := lexGuardStep([]rune(runes), i)
	if wantErr {
		if err == nil {
			t.Fatalf("lexGuardStep(%q,%d) expected error", runes, i)
		}
		return
	}
	if err != nil {
		t.Fatalf("lexGuardStep(%q,%d) unexpected error: %v", runes, i, err)
	}
	if emit != wantEmit {
		t.Errorf("emit = %v, want %v", emit, wantEmit)
	}
	if next != wantNext {
		t.Errorf("next = %d, want %d", next, wantNext)
	}
	if wantEmit && (tok.kind != wantKind || tok.text != wantText) {
		t.Errorf("tok = {%d %q}, want {%d %q}", tok.kind, tok.text, wantKind, wantText)
	}
}
