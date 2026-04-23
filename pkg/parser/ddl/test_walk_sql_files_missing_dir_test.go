//ff:func feature=ddl type=test control=sequence
//ff:what walkSQLFiles — 존재하지 않는 디렉토리는 handler 미호출 + 1개 진단 반환

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestWalkSQLFiles_MissingDir(t *testing.T) {
	called := false
	diags := walkSQLFiles("/nonexistent/ddl/walksql/xyz123", func(path string, data []byte) []diagnostic.Diagnostic {
		called = true
		return nil
	})
	if called {
		t.Fatalf("handler must not be called when dir read fails")
	}
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1", len(diags))
	}
	d := diags[0]
	if d.File != "/nonexistent/ddl/walksql/xyz123" {
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
