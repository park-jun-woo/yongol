//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what runStructSetCase — map[string]struct{} 비교 테스트 공통 헬퍼

package hurl_openapi

import "testing"

func runStructSetCase(t *testing.T, got, want map[string]struct{}) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d entries, want %d; got=%v", len(got), len(want), got)
	}
	for k := range want {
		if _, ok := got[k]; !ok {
			t.Errorf("missing key %q", k)
		}
	}
}
