//ff:func feature=contract type=test control=sequence
//ff:what test: TestExpandFieldList — nil 리스트, 다중 이름 그룹, 익명 반환 타입 전개 분기 검증
package contract

import (
	"testing"
)

func TestExpandFieldListUnnamedResults(t *testing.T) {
	fset, fd := fieldListFromFunc(t, "func F() (string, error) {}")
	got := expandFieldList(fset, fd.Type.Results, false)
	if len(got) != 2 {
		t.Fatalf("got %d results (%v) want 2", len(got), got)
	}
	if got[0].Type != "string" || got[0].Name != "" {
		t.Errorf("result[0]: got %+v want unnamed string", got[0])
	}
	if got[1].Type != "error" || got[1].Name != "" {
		t.Errorf("result[1]: got %+v want unnamed error", got[1])
	}
}
