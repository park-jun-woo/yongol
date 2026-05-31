//ff:func feature=manifest type=test control=sequence
//ff:what drainPendingHints — pending 힌트에 table/column 컨텍스트 부착 후 방출
package ddl

import (
	"testing"
)

func TestDrainPendingHints(t *testing.T) {
	t.Run("attaches column and table", func(t *testing.T) {
		pending := []*HintComment{{Tag: "cast"}, {Tag: "rename"}}
		out, remaining := drainPendingHints(pending, "users", "email VARCHAR(255)", nil)
		if remaining != nil {
			t.Errorf("remaining = %v, want nil", remaining)
		}
		if len(out) != 2 {
			t.Fatalf("out len = %d, want 2", len(out))
		}
		for _, h := range out {
			if h.TableCtx != "users" || h.ColumnCtx != "email" {
				t.Errorf("hint ctx = %+v", h)
			}
		}
	})
	t.Run("no resolvable column keeps column empty", func(t *testing.T) {
		pending := []*HintComment{{Tag: "data_migration"}}
		out, _ := drainPendingHints(pending, "users", "", nil)
		if len(out) != 1 || out[0].ColumnCtx != "" {
			t.Errorf("out = %+v", out)
		}
		if out[0].TableCtx != "users" {
			t.Errorf("TableCtx = %q, want users", out[0].TableCtx)
		}
	})
}
