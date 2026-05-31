//ff:func feature=ssac-parse type=test control=sequence
//ff:what parseResult/parseGuard/parseCRUD(no-result) 파싱 검증
package ssac

import (
	"testing"
)

func TestParseGuard(t *testing.T) {
	t.Run("target and message", func(t *testing.T) {
		seq := parseGuard("empty", `course "not found"`)
		if seq.Type != "empty" || seq.Target != "course" || seq.Message != "not found" {
			t.Errorf("seq = %+v", seq)
		}
		if seq.ErrStatus != 0 {
			t.Errorf("ErrStatus = %d, want 0", seq.ErrStatus)
		}
	})
	t.Run("with explicit status remainder", func(t *testing.T) {
		seq := parseGuard("exists", `dup "conflict" 422`)
		if seq.ErrStatus != 422 {
			t.Errorf("ErrStatus = %d, want 422", seq.ErrStatus)
		}
	})
	t.Run("non-numeric remainder ignored", func(t *testing.T) {
		seq := parseGuard("empty", `x "msg" notanumber`)
		if seq.ErrStatus != 0 {
			t.Errorf("ErrStatus = %d, want 0", seq.ErrStatus)
		}
	})
}
