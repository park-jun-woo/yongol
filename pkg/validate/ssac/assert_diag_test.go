//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ssac-structural
//ff:what assertDiag — diags 에 prefix 포함 진단 존재 확인 (없으면 t.Error)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// assertDiag fails t if none of diags contains the given prefix.
func assertDiag(t *testing.T, diags []diagnostic.Diagnostic, prefix string) {
	t.Helper()
	for _, d := range diags {
		if strings.Contains(d.Message, prefix) {
			return
		}
	}
	t.Errorf("expected diagnostic with %s, got %d diags", prefix, len(diags))
}
