//ff:func feature=manifest type=test control=sequence
//ff:what applyTableCheckEnum — 테이블 레벨 CHECK enum을 해당 Column에 부착 / 미존재 컬럼 no-op
package ddl

import (
	"testing"
)

func TestApplyTableCheckEnum(t *testing.T) {
	t.Run("attaches to existing column", func(t *testing.T) {
		tb := &Table{Columns: map[string]Column{"status": {}}}
		applyTableCheckEnum("CHECK (status IN ('open','closed'))", tb)
		got := tb.Columns["status"]
		if len(got.CheckEnum) != 2 {
			t.Errorf("CheckEnum = %v, want 2 values", got.CheckEnum)
		}
	})
	t.Run("unknown column is no-op", func(t *testing.T) {
		tb := &Table{Columns: map[string]Column{"status": {}}}
		applyTableCheckEnum("CHECK (missing IN ('a'))", tb)
		if len(tb.Columns["status"].CheckEnum) != 0 {
			t.Errorf("status should be untouched")
		}
		if _, ok := tb.Columns["missing"]; ok {
			t.Errorf("missing column should not be created")
		}
	})
	t.Run("no enum is no-op", func(t *testing.T) {
		tb := &Table{Columns: map[string]Column{"id": {}}}
		applyTableCheckEnum("CHECK (id > 0)", tb)
		if len(tb.Columns["id"].CheckEnum) != 0 {
			t.Errorf("id should be untouched")
		}
	})
}
