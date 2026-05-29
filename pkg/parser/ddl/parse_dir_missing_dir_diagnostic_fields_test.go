//ff:func feature=manifest type=test control=sequence
//ff:what ParseDir — 존재하지 않는 디렉토리에서 Phase/Level/File 진단 필드 완결성

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseDir_MissingDir_DiagnosticFields(t *testing.T) {
	_, diags := ParseDir("/nonexistent/ddl/xyz123")
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1", len(diags))
	}
	d := diags[0]
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want LevelError", d.Level)
	}
	if d.File != "/nonexistent/ddl/xyz123" {
		t.Errorf("File = %q, want dir", d.File)
	}
}
