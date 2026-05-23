//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=funcspec-structural
//ff:what runXff41FuncForbiddenImport — TestXff41FuncForbiddenImport table-driven 개별 케이스 검증

package funcspec

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runXff41FuncForbiddenImport(t *testing.T, c TestXff41FuncForbiddenImportCase) {
	t.Helper()
	fs := &yongol.Fullstack{ProjectFuncSpecs: c.specs}
	diags := xff41FuncForbiddenImport(fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[XFF-41]") {
			t.Errorf("message should contain [XFF-41], got %q", d.Message)
		}
	}
}
