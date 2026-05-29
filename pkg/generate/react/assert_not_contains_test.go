//ff:func feature=gen-react type=test control=sequence
//ff:what 테스트 헬퍼 — 문자열 미포함 여부 검증

package react

import (
	"strings"
	"testing"
)

func assertNotContains(t *testing.T, content, unwanted string) {
	t.Helper()
	if strings.Contains(content, unwanted) {
		t.Errorf("output should not contain %q but does:\n%s", unwanted, content)
	}
}
