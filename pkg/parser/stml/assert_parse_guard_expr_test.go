//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardExpr 단일 케이스 kind/op/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardExpr lexes input, runs parseGuardExpr, and asserts the kind
// (and op for binary), or the expected error.
func assertParseGuardExpr(t *testing.T, input string, wantErr bool, wantKind GuardKind, wantOp string) {
	t.Helper()
	toks, err := lexGuard(input)
	if err != nil {
		t.Fatalf("lexGuard(%q) error: %v", input, err)
	}
	p := &guardParser{toks: toks}
	expr, err := p.parseGuardExpr()
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardExpr(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardExpr(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != wantKind {
		t.Errorf("kind = %d, want %d", expr.Kind, wantKind)
	}
	if wantKind == GuardBinary && expr.Op != wantOp {
		t.Errorf("op = %q, want %q", expr.Op, wantOp)
	}
}
