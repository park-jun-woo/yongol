//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-infra
//ff:what runBoolSetCase — map[string]bool 크기 및 포함 키 검증 공통 헬퍼

package manifest

import "testing"

func runBoolSetCase(t *testing.T, got map[string]bool, wantN int, wantIn []string) {
	t.Helper()
	if len(got) != wantN {
		t.Fatalf("got %d, want %d", len(got), wantN)
	}
	for _, k := range wantIn {
		if !got[k] {
			t.Errorf("missing %q", k)
		}
	}
}
