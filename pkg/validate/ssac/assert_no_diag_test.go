//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ssac-structural
//ff:what assertNoDiag — diags 에 prefix 포함 진단 부존재 확인 (있으면 t.Error)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// assertNoDiag fails t if any of diags contains the given prefix.
func assertNoDiag(t *testing.T, diags []diagnostic.Diagnostic, prefix string) {
	t.Helper()
	for _, d := range diags {
		if strings.Contains(d.Message, prefix) {
			t.Errorf("unexpected diagnostic %s: %s", prefix, d.Message)
			return
		}
	}
}
