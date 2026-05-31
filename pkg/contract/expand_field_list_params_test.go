//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExpandFieldList — nil 리스트, 다중 이름 그룹, 익명 반환 타입 전개 분기 검증
package contract

import (
	"testing"
)

func TestExpandFieldListParams(t *testing.T) {
	fset, fd := fieldListFromFunc(t, "func F(a, b int, name string) {}")
	got := expandFieldList(fset, fd.Type.Params, true)
	want := []FuncParam{
		{Name: "a", Type: "int"},
		{Name: "b", Type: "int"},
		{Name: "name", Type: "string"},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d params (%v) want %d", len(got), got, len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("param[%d]: got %+v want %+v", i, got[i], want[i])
		}
	}
}
