//ff:func feature=gen-hurl type=test-helper control=iteration dimension=1
//ff:what assertWarningContains — warnings 에 sub 포함 여부 검증

package hurl

import (
	"strings"
	"testing"
)

// assertWarningContains fails the test when no warning contains sub.
func assertWarningContains(t *testing.T, warns []string, sub string) {
	t.Helper()
	for _, w := range warns {
		if strings.Contains(w, sub) {
			return
		}
	}
	t.Errorf("expected warning substring %q; got %v", sub, warns)
}
