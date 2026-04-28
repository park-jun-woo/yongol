//ff:func feature=migration type=test-helper control=sequence
//ff:what inspectCreateIndexForMethod — op 이 CreateIndex 면 Method 검증하고 true 반환

package migration

import "testing"

// inspectCreateIndexForMethod checks whether op is a CreateIndex. When it
// is, the helper also asserts the Index.Method equals wantMethod (fatal on
// mismatch) and returns true; otherwise false.
func inspectCreateIndexForMethod(t *testing.T, op Operation, wantMethod string) bool {
	t.Helper()
	ci, ok := op.(CreateIndex)
	if !ok {
		return false
	}
	if ci.Index.Method != wantMethod {
		t.Errorf("CreateIndex.Method = %q, want %q", ci.Index.Method, wantMethod)
	}
	return true
}
