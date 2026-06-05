//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardAtom 단일 케이스 분기/경로/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardAtom lexes input, runs parseGuardAtom, and asserts the
// resulting kind and ref path, or the expected error.
func assertParseGuardAtom(t *testing.T, input string, wantErr bool, wantKind GuardKind, wantPath string) {
	t.Helper()
	toks, err := lexGuard(input)
	if err != nil {
		t.Fatalf("lexGuard(%q) error: %v", input, err)
	}
	p := &guardParser{toks: toks}
	expr, err := p.parseGuardAtom()
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardAtom(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardAtom(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != wantKind {
		t.Errorf("kind = %d, want %d", expr.Kind, wantKind)
	}
	path := expr.Ref.Path()
	if wantKind == GuardGroup {
		path = expr.Operand.Ref.Path()
	}
	if path != wantPath {
		t.Errorf("path = %q, want %q", path, wantPath)
	}
}
