//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestNilCheckExpr — NotNull 여부에 따른 None 검사 표현식 검증

package types

import "testing"

func TestNilCheckExpr(t *testing.T) {
	if got := nilCheckExpr(true); got != "" {
		t.Errorf("nilCheckExpr(true) = %q, want empty", got)
	}
	if got := nilCheckExpr(false); got != "{var} is None" {
		t.Errorf("nilCheckExpr(false) = %q", got)
	}
}
