//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardTerm 단일 케이스 kind/부정 op/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardTerm lexes input, runs parseGuardTerm, and asserts the kind
// (and "!" op for unary), or the expected error.
func assertParseGuardTerm(t *testing.T, input string, wantErr bool, wantKind GuardKind) {
	t.Helper()
	toks, err := lexGuard(input)
	if err != nil {
		t.Fatalf("lexGuard(%q) error: %v", input, err)
	}
	p := &guardParser{toks: toks}
	expr, err := p.parseGuardTerm()
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardTerm(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardTerm(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != wantKind {
		t.Errorf("kind = %d, want %d", expr.Kind, wantKind)
	}
	if wantKind == GuardUnary && expr.Op != "!" {
		t.Errorf("op = %q, want %q", expr.Op, "!")
	}
}
