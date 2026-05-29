//ff:func feature=ssac-parse type=parser control=sequence
//ff:what assertEqual 헬퍼 — 문자열 값 비교 실패 시 테스트 에러 출력

package ssac

import "testing"

func assertEqual(t *testing.T, name, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %q, want %q", name, got, want)
	}
}
