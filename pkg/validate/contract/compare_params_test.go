//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareParams — 파라미터 타입 불일치를 인덱스 단위 diff 로 반환 검증
package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareParams(t *testing.T) {
	t.Run("type mismatch", func(t *testing.T) {
		exp := []contract.FuncParam{{Name: "a", Type: "int"}, {Name: "b", Type: "string"}}
		act := []contract.FuncParam{{Name: "x", Type: "int"}, {Name: "y", Type: "int64"}}
		diffs := compareParams(exp, act)
		if len(diffs) != 1 {
			t.Fatalf("expected 1 diff, got %v", diffs)
		}
		if !strings.Contains(diffs[0], "param #2 type") {
			t.Errorf("diff = %q", diffs[0])
		}
	})

	t.Run("names ignored", func(t *testing.T) {
		exp := []contract.FuncParam{{Name: "a", Type: "int"}}
		act := []contract.FuncParam{{Name: "z", Type: "int"}}
		if diffs := compareParams(exp, act); len(diffs) != 0 {
			t.Errorf("expected no diffs for name-only difference, got %v", diffs)
		}
	})
}
