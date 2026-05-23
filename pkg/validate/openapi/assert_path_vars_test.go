//ff:func feature=validate type=test-helper control=iteration dimension=1
//ff:what assertPathVars — collectPathVars 결과 키 매칭 검증 헬퍼

package openapi

import "testing"

func assertPathVars(t *testing.T, path string, want map[string]bool) {
	t.Helper()
	got := collectPathVars(path)
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d: %v", len(got), len(want), got)
	}
	for k := range want {
		if !got[k] {
			t.Errorf("missing key %q in result: %v", k, got)
		}
	}
}
