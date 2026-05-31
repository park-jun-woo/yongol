//ff:func feature=manifest type=test control=sequence
//ff:what appendPendingHint — 유효 힌트는 누적, 비힌트는 pending 불변
package ddl

import (
	"testing"
)

func TestAppendPendingHint(t *testing.T) {
	t.Run("valid hint appended", func(t *testing.T) {
		got := appendPendingHint("-- @backfill", "/x.sql", 2, "users", nil)
		if len(got) != 1 {
			t.Fatalf("pending len = %d, want 1", len(got))
		}
		if got[0].Tag != "backfill" || got[0].TableCtx != "users" {
			t.Errorf("hint = %+v", got[0])
		}
	})
	t.Run("non-hint leaves pending unchanged", func(t *testing.T) {
		base := []*HintComment{{Tag: "rename"}}
		got := appendPendingHint("-- not a hint", "", 1, "", base)
		if len(got) != 1 {
			t.Errorf("pending len = %d, want 1 (unchanged)", len(got))
		}
	})
}
