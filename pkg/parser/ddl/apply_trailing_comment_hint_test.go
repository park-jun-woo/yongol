//ff:func feature=manifest type=test control=sequence
//ff:what applyTrailingCommentHint — 인라인 `-- @...` 주석을 HintComment로 파싱해 out에 추가
package ddl

import (
	"testing"
)

func TestApplyTrailingCommentHint(t *testing.T) {
	t.Run("hint comment appended with column ctx", func(t *testing.T) {
		out := applyTrailingCommentHint("@cast type=int", "amount NUMERIC", "/x.sql", 5, "orders", nil)
		if len(out) != 1 {
			t.Fatalf("out len = %d, want 1", len(out))
		}
		h := out[0]
		if h.Tag != "cast" || h.ColumnCtx != "amount" || h.TableCtx != "orders" {
			t.Errorf("hint = %+v", h)
		}
	})
	t.Run("non-hint comment not appended", func(t *testing.T) {
		out := applyTrailingCommentHint("just a note", "id BIGINT", "/x.sql", 1, "t", nil)
		if len(out) != 0 {
			t.Errorf("out len = %d, want 0", len(out))
		}
	})
}
