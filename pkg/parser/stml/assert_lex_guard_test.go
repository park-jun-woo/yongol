//ff:func feature=stml-parse type=test-helper control=iteration dimension=1
//ff:what lexGuard 단일 케이스 토큰 슬라이스/에러 검증 헬퍼

package stml

import "testing"

// assertLexGuard runs lexGuard on input and asserts the produced token slice
// matches want kind/text per index, or the expected error.
func assertLexGuard(t *testing.T, input string, want []guardToken, wantErr bool) {
	t.Helper()
	got, err := lexGuard(input)
	if wantErr {
		if err == nil {
			t.Fatalf("lexGuard(%q) expected error, got nil", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("lexGuard(%q) unexpected error: %v", input, err)
	}
	if len(got) != len(want) {
		t.Fatalf("lexGuard(%q) token count = %d, want %d (got %+v)", input, len(got), len(want), got)
	}
	for i := range want {
		if got[i].kind != want[i].kind || got[i].text != want[i].text {
			t.Errorf("token[%d] = {%d %q}, want {%d %q}", i, got[i].kind, got[i].text, want[i].kind, want[i].text)
		}
	}
}
