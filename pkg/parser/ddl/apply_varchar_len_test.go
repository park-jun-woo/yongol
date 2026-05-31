//ff:func feature=manifest type=test control=sequence
//ff:what applyVarcharLen — VARCHAR(N) 길이 설정 / 비VARCHAR 무시
package ddl

import (
	"testing"
)

func TestApplyVarcharLen(t *testing.T) {
	t.Run("sets length", func(t *testing.T) {
		col := &Column{}
		applyVarcharLen(col, "VARCHAR(255)")
		if col.VarcharLen != 255 {
			t.Errorf("VarcharLen = %d, want 255", col.VarcharLen)
		}
	})
	t.Run("no length leaves zero", func(t *testing.T) {
		col := &Column{}
		applyVarcharLen(col, "BIGINT")
		if col.VarcharLen != 0 {
			t.Errorf("VarcharLen = %d, want 0", col.VarcharLen)
		}
	})
}
