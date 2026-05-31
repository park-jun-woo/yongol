//ff:func feature=contract type=test control=sequence
//ff:what test: TestExpandFieldList — nil 리스트, 다중 이름 그룹, 익명 반환 타입 전개 분기 검증
package contract

import (
	"go/token"
	"testing"
)

func TestExpandFieldListNil(t *testing.T) {
	fset := token.NewFileSet()
	if got := expandFieldList(fset, nil, false); got != nil {
		t.Errorf("nil list: got %v want nil", got)
	}
}
