//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=funcspec-structural
//ff:what runCollectForbiddenImports — TestCollectForbiddenImports table-driven 개별 케이스 검증

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runCollectForbiddenImports(t *testing.T, c TestCollectForbiddenImportsCase) {
	t.Helper()
	diags := collectForbiddenImports(c.pkg, c.funcName, c.line, c.imports)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if d.Phase != diagnostic.PhaseValidate {
			t.Errorf("expected PhaseValidate, got %q", d.Phase)
		}
		wantFile := c.pkg + "/" + c.funcName + ".go"
		if d.File != wantFile {
			t.Errorf("File = %q, want %q", d.File, wantFile)
		}
		if d.Line != c.line {
			t.Errorf("Line = %d, want %d", d.Line, c.line)
		}
	}
}
