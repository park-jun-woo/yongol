//ff:func feature=manifest type=test control=sequence
//ff:what ParseTables — 존재하지 않는 디렉토리에서 Diagnostic 필드(File/Phase/Level/Message) 완결성 검증

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseTables_MissingDir_DiagnosticFields(t *testing.T) {
	_, diags := ParseTables("/nonexistent/ddl/dir/abcxyz")
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1: %v", len(diags), diags)
	}
	d := diags[0]
	if d.File != "/nonexistent/ddl/dir/abcxyz" {
		t.Errorf("File = %q, want dir path", d.File)
	}
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want LevelError", d.Level)
	}
	if d.Message == "" {
		t.Errorf("Message empty")
	}
}
