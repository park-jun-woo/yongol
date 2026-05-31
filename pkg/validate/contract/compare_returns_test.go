//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareReturns — 반환 타입 불일치를 인덱스 단위 diff 로 반환 검증
package contract

import (
	"strings"
	"testing"
)

func TestCompareReturns(t *testing.T) {
	t.Run("mismatch", func(t *testing.T) {
		diffs := compareReturns([]string{"int", "error"}, []string{"int64", "error"})
		if len(diffs) != 1 {
			t.Fatalf("expected 1 diff, got %v", diffs)
		}
		if !strings.Contains(diffs[0], "return #1 type") {
			t.Errorf("diff = %q", diffs[0])
		}
	})

	t.Run("identical", func(t *testing.T) {
		if diffs := compareReturns([]string{"error"}, []string{"error"}); len(diffs) != 0 {
			t.Errorf("expected no diffs, got %v", diffs)
		}
	})
}
