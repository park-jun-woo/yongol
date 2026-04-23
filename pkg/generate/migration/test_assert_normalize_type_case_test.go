//ff:func feature=migration type=test-helper control=sequence
//ff:what assertNormalizeTypeCase — 단일 normalizeTypeCase 기대값과 NormalizeType 결과 비교
package migration

import "testing"

func assertNormalizeTypeCase(t *testing.T, c normalizeTypeCase) {
	t.Helper()
	got, isSerial := NormalizeType(c.in)
	if got.Base != c.wantBase {
		t.Errorf("Base: got %q, want %q", got.Base, c.wantBase)
	}
	if got.Length != c.wantLen {
		t.Errorf("Length: got %d, want %d", got.Length, c.wantLen)
	}
	if got.Precision != c.wantPrec {
		t.Errorf("Precision: got %d, want %d", got.Precision, c.wantPrec)
	}
	if got.Scale != c.wantScale {
		t.Errorf("Scale: got %d, want %d", got.Scale, c.wantScale)
	}
	if got.Array != c.wantArray {
		t.Errorf("Array: got %v, want %v", got.Array, c.wantArray)
	}
	if isSerial != c.wantSerial {
		t.Errorf("isSerial: got %v, want %v", isSerial, c.wantSerial)
	}
}
