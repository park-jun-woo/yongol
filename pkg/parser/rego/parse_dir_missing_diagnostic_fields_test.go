//ff:func feature=policy type=test control=sequence
//ff:what ParseDir — 존재하지 않는 디렉토리에서 Phase/Level/File 완결성

package rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseDir_Missing_DiagnosticFields(t *testing.T) {
	_, diags := ParseDir("/nonexistent/rego/dir")
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
	if d.File != "/nonexistent/rego/dir" {
		t.Errorf("File = %q", d.File)
	}
}
