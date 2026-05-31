//ff:func feature=funcspec type=test-helper control=iteration dimension=1
//ff:what assertExtractReturnTypes — extractReturnTypes 결과가 want 와 일치하는지 검증 헬퍼
package funcspec

import "testing"

// assertExtractReturnTypes parses src and asserts extractReturnTypes on its
// first func equals want.
func assertExtractReturnTypes(t *testing.T, src string, want []string) {
	t.Helper()
	fset, f := parseDeclT(t, src)
	got := extractReturnTypes(fset, firstFunc(f))
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
