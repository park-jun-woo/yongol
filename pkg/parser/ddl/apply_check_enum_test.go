//ff:func feature=manifest type=test control=sequence
//ff:what applyCheckEnum — CHECK IN (...) 값을 Column.CheckEnum 에 적용
package ddl

import (
	"testing"
)

func TestApplyCheckEnum(t *testing.T) {
	t.Run("captures enum values", func(t *testing.T) {
		col := &Column{}
		applyCheckEnum("role VARCHAR(20) CHECK (role IN ('member','admin'))", col)
		if len(col.CheckEnum) != 2 {
			t.Fatalf("CheckEnum len = %d, want 2 (%v)", len(col.CheckEnum), col.CheckEnum)
		}
	})
	t.Run("no enum leaves nil", func(t *testing.T) {
		col := &Column{}
		applyCheckEnum("id BIGINT NOT NULL", col)
		if col.CheckEnum != nil {
			t.Errorf("CheckEnum = %v, want nil", col.CheckEnum)
		}
	})
}
