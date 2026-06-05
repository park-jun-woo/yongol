//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardCompare 단일 케이스 op/value/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardCompare lexes input, parses a ref then a compare, and asserts
// the op and value, or the expected error.
func assertParseGuardCompare(t *testing.T, input string, wantErr bool, wantOp, wantValue string) {
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
	expr, err := p.parseGuardCompare(ref)
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardCompare(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardCompare(%q) unexpected error: %v", input, err)
	}
	if expr.Kind != GuardCompare {
		t.Errorf("kind = %d, want GuardCompare", expr.Kind)
	}
	if expr.Op != wantOp {
		t.Errorf("op = %q, want %q", expr.Op, wantOp)
	}
	if expr.Value != wantValue {
		t.Errorf("value = %q, want %q", expr.Value, wantValue)
	}
}
