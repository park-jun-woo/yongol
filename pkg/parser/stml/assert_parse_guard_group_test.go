//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardGroup 단일 케이스 그룹 내부 경로/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardGroup lexes input, runs parseGuardGroup, and asserts the inner
// operand path, or the expected error.
func assertParseGuardGroup(t *testing.T, input string, wantErr bool, wantPath string) {
	t.Helper()
	toks, err := lexGuard(input)
	if err != nil {
		t.Fatalf("lexGuard(%q) error: %v", input, err)
	}
	p := &guardParser{toks: toks}
	expr, err := p.parseGuardGroup()
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardGroup(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardGroup(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != GuardGroup {
		t.Errorf("kind = %d, want GuardGroup", expr.Kind)
	}
	if expr.Operand.Ref.Path() != wantPath {
		t.Errorf("inner path = %q, want %q", expr.Operand.Ref.Path(), wantPath)
	}
}
