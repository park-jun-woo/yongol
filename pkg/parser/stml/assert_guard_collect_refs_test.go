//ff:func feature=stml-parse type=test-helper control=iteration dimension=1
//ff:what 가드 조건을 파싱해 CollectRefs 결과 경로가 기대 목록과 일치하는지 검증한다

package stml

import "testing"

// assertGuardCollectRefs parses condition and asserts that CollectRefs yields
// refs whose dotted paths equal want, in order.
func assertGuardCollectRefs(t *testing.T, condition string, want []string) {
	t.Helper()
	expr, err := ParseGuard(condition)
	if err != nil {
		t.Fatalf("ParseGuard(%q) error: %v", condition, err)
	}
	refs := expr.CollectRefs()
	if len(refs) != len(want) {
		t.Fatalf("CollectRefs len = %d, want %d (%+v)", len(refs), len(want), refs)
	}
	for i, w := range want {
		if refs[i].Path() != w {
			t.Errorf("refs[%d].Path() = %q, want %q", i, refs[i].Path(), w)
		}
	}
}
