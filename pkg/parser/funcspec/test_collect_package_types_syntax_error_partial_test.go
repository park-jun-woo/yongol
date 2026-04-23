//ff:func feature=funcspec type=test control=sequence
//ff:what collectPackageTypes — 한 파일 syntax error 에도 다른 파일의 struct 는 수집 (partial success)

package funcspec

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCollectPackageTypesSyntaxErrorPartial — a file with a syntax error must be reported as
// exactly one Diagnostic while other valid files in the same directory are still collected (partial success).
func TestCollectPackageTypesSyntaxErrorPartial(t *testing.T) {
	dir := t.TempDir()

	ok := `package sample

type OkRequest struct {
	Name string
}
`
	bad := `package sample

this is not valid go at all !!!
`
	if err := os.WriteFile(filepath.Join(dir, "ok.go"), []byte(ok), 0644); err != nil {
		t.Fatal(err)
	}
	badPath := filepath.Join(dir, "bad.go")
	if err := os.WriteFile(badPath, []byte(bad), 0644); err != nil {
		t.Fatal(err)
	}

	result, diags := collectPackageTypes(dir)
	if len(diags) != 1 {
		t.Fatalf("diags = %d, want 1: %v", len(diags), diags)
	}
	d := diags[0]
	if d.File != badPath {
		t.Errorf("diag.File = %q, want %q", d.File, badPath)
	}
	if d.Phase != "parse" {
		t.Errorf("diag.Phase = %q, want parse", d.Phase)
	}
	if d.Level != "ERROR" {
		t.Errorf("diag.Level = %q, want ERROR", d.Level)
	}
	if d.Line <= 0 {
		t.Errorf("diag.Line = %d, want > 0 (extractGoParseErrorLine)", d.Line)
	}
	if !strings.Contains(d.Message, "Go parse failed") {
		t.Errorf("diag.Message = %q, want 'Go parse failed' prefix", d.Message)
	}
	// partial success: OkRequest from ok.go must still be collected.
	if _, found := result["OkRequest"]; !found {
		t.Errorf("OkRequest should still be collected despite sibling syntax error; keys=%v", keysOf(result))
	}
}
