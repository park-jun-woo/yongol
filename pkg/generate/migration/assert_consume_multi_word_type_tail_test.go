//ff:func feature=migration type=test-helper control=sequence
//ff:what assertConsumeMultiWordTypeTail — consumeMultiWordTypeTail 의 idx/parts 결과 검증 헬퍼
package migration

import (
	"reflect"
	"testing"
)

// assertConsumeMultiWordTypeTail copies startPart and asserts the resulting
// index and parts from consumeMultiWordTypeTail.
func assertConsumeMultiWordTypeTail(t *testing.T, toks []string, i int, startPart, wantParts []string, wantIdx int) {
	t.Helper()
	parts := make([]string, len(startPart))
	copy(parts, startPart)
	gotIdx := consumeMultiWordTypeTail(nil, toks, i, &parts)
	if gotIdx != wantIdx {
		t.Errorf("idx = %d, want %d", gotIdx, wantIdx)
	}
	if !reflect.DeepEqual(parts, wantParts) {
		t.Errorf("parts = %#v, want %#v", parts, wantParts)
	}
}
