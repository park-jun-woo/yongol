//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardLifecycle 단일 케이스 lifecycle/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardLifecycle lexes input, parses a ref then a lifecycle, and
// asserts the lifecycle value, or the expected error.
func assertParseGuardLifecycle(t *testing.T, input string, wantErr bool, wantLifecycle string) {
	t.Helper()
	toks, err := lexGuard(input)
	if err != nil {
		t.Fatalf("lexGuard(%q) error: %v", input, err)
	}
	p := &guardParser{toks: toks}
	ref, err := p.parseGuardRef()
	if err != nil {
		t.Fatalf("parseGuardRef(%q) error: %v", input, err)
	}
	expr, err := p.parseGuardLifecycle(ref)
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardLifecycle(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardLifecycle(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != GuardLifecycle {
		t.Errorf("kind = %d, want GuardLifecycle", expr.Kind)
	}
	if expr.Lifecycle != wantLifecycle {
		t.Errorf("lifecycle = %q, want %q", expr.Lifecycle, wantLifecycle)
	}
}
