//ff:func feature=contract type=test control=sequence
//ff:what test: TestWalkForSymbols — body 2-pass walk 으로 SqlcQueries·CallTargets·DDLFields 정렬 분류 검증
package contract

import (
	"testing"
)

func TestWalkForSymbolsEmpty(t *testing.T) {
	fset, body := bodyFromFunc(t, "x := 1\n_ = x\n")
	sym := walkForSymbols(fset, body)
	if sym.SqlcQueries != nil || sym.CallTargets != nil || sym.DDLFields != nil {
		t.Errorf("expected all nil for body with no external symbols, got %+v", sym)
	}
}
