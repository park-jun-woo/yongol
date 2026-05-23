//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what runBoolMapCase — map[string]bool 비교 테스트 공통 헬퍼

package hurl_openapi

import "testing"

func runBoolMapCase(t *testing.T, got, want map[string]bool) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d entries, want %d; got=%v", len(got), len(want), got)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("key %q: got %v, want %v", k, got[k], v)
		}
	}
}
