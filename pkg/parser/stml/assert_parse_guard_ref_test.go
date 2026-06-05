//ff:func feature=stml-parse type=test-helper control=sequence
//ff:what parseGuardRef 단일 케이스 model/field/에러 검증 헬퍼

package stml

import "testing"

// assertParseGuardRef lexes input, runs parseGuardRef, and asserts the model and
// field, or the expected error.
func assertParseGuardRef(t *testing.T, input string, wantErr bool, wantModel, wantField string) {
	t.Helper()
	toks, err := lexGuard(input)
	if err != nil {
		t.Fatalf("lexGuard(%q) error: %v", input, err)
	}
	p := &guardParser{toks: toks}
	ref, err := p.parseGuardRef()
	if wantErr {
		if err == nil {
			t.Fatalf("parseGuardRef(%q) expected error", input)
		}
		return
	}
	if err != nil {
		t.Fatalf("parseGuardRef(%q) unexpected error: %v", input, err)
	}
	if ref.Model != wantModel {
		t.Errorf("model = %q, want %q", ref.Model, wantModel)
	}
	if ref.Field != wantField {
		t.Errorf("field = %q, want %q", ref.Field, wantField)
	}
}
