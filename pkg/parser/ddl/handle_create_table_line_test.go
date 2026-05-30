//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what handleCreateTableLine — pending 힌트를 block-above 로 소비하고 table ctx 갱신

package ddl

import "testing"

func TestHandleCreateTableLine(t *testing.T) {
	pending := []*HintComment{{Tag: "rename"}, {Tag: "cast"}}
	ctx, out, remaining := handleCreateTableLine("CREATE TABLE orders (", pending, nil)
	if ctx != "orders" {
		t.Errorf("ctx = %q, want orders", ctx)
	}
	if remaining != nil {
		t.Errorf("remaining = %v, want nil", remaining)
	}
	if len(out) != 2 {
		t.Fatalf("out len = %d, want 2", len(out))
	}
	for _, h := range out {
		if h.TableCtx != "orders" || !h.BlockAbove {
			t.Errorf("hint = %+v, want TableCtx=orders BlockAbove=true", h)
		}
	}
}
