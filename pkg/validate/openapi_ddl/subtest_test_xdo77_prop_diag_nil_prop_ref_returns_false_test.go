//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagNilPropRefReturnsFalse — nil propRef returns false 서브테스트
package openapi_ddl

import (
	"testing"
)

func subtestTestXdo77PropDiagNilPropRefReturnsFalse(t *testing.T) {

	_, ok := xdo77PropDiag(xdo77FS(), "User", "users", "id", nil, xdo77Cols())
	if ok {
		t.Error("expected false for nil propRef")
	}

}
