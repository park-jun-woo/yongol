//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what ParseGuard 단일 케이스 kind/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuard runs ParseGuard on input and asserts the expression kind, or
// the expected error.
func assertParseGuard(t *testing.T, input string, wantErr bool, wantKind GuardKind) {
	t.Helper()
	expr, err := ParseGuard(input)
	if wantErr {
		if err == nil {
			t.Fatalf("ParseGuard(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("ParseGuard(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != wantKind {
		t.Errorf("kind = %d, want %d", expr.Kind, wantKind)
	}
}
