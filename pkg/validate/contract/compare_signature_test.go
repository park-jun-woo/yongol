//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareSignature — 두 FuncSignature 비교 시 name/arity/type/err 차이 종합 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareSignature(t *testing.T) {
	base := contract.FuncSignature{
		Name:    "Do",
		Params:  []contract.FuncParam{{Type: "int"}},
		Returns: []string{"error"},
		HasErr:  true,
	}

	t.Run("identical → no diffs", func(t *testing.T) {
		if diffs := compareSignature(base, base); len(diffs) != 0 {
			t.Errorf("expected no diffs, got %v", diffs)
		}
	})

	t.Run("name + type + return-arity + err", func(t *testing.T) {
		act := contract.FuncSignature{
			Name:    "Done",
			Params:  []contract.FuncParam{{Type: "string"}},
			Returns: []string{},
			HasErr:  false,
		}
		diffs := compareSignature(base, act)
		joined := strings.Join(diffs, "|")
		for _, want := range []string{"function name", "param #1 type", "return arity", "error-return"} {
			if !strings.Contains(joined, want) {
				t.Errorf("missing %q in diffs: %v", want, diffs)
			}
		}
	})
}
