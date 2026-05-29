//ff:func feature=gen-react type=test control=sequence
//ff:what 테스트 헬퍼 — 문자열 포함 여부 검증

package react

import (
	"strings"
	"testing"
)

func assertContains(t *testing.T, content, want string) {
	t.Helper()
	if !strings.Contains(content, want) {
		t.Errorf("output missing expected substring:\n  want: %q\n  got:\n%s", want, content)
	}
}
