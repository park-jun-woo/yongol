//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=funcspec-structural
//ff:what runXff40FuncBodyTodo — TestXff40FuncBodyTodo table-driven 개별 케이스 검증

package funcspec

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runXff40FuncBodyTodo(t *testing.T, c TestXff40FuncBodyTodoCase) {
	t.Helper()
	fs := &yongol.Fullstack{ProjectFuncSpecs: c.specs}
	diags := xff40FuncBodyTodo(fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[XFF-40]") {
			t.Errorf("message should contain [XFF-40], got %q", d.Message)
		}
	}
}
