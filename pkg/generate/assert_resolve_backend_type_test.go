//ff:func feature=generate type=test-helper control=sequence
//ff:what assertResolveBackendType — ResolveBackendType 결과(에러 기대 / 값) 검증 헬퍼
package generate

import "testing"

// assertResolveBackendType asserts ResolveBackendType(lang, fw) matches the
// expected backend type or error expectation.
func assertResolveBackendType(t *testing.T, lang, fw string, want BackendType, wantErr bool) {
	t.Helper()
	got, err := ResolveBackendType(lang, fw)
	if wantErr {
		if err == nil {
			t.Errorf("ResolveBackendType(%q,%q) expected error", lang, fw)
		}
		return
	}
	if err != nil || got != want {
		t.Errorf("ResolveBackendType(%q,%q) = (%q,%v), want %q", lang, fw, got, err, want)
	}
}
