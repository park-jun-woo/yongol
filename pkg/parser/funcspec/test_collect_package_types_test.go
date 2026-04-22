//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what validates Diagnostic propagation behaviour of collectPackageTypes / fillMissingFields

package funcspec

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCollectPackageTypesNormal — a valid Go file must collect structs with zero diagnostics.
func TestCollectPackageTypesNormal(t *testing.T) {
	dir := t.TempDir()
	src := `package sample

type FooRequest struct {
	Name string ` + "`json:\"name\"`" + `
}
type FooResponse struct {
	Id int ` + "`json:\"id\"`" + `
}
`
	if err := os.WriteFile(filepath.Join(dir, "foo.go"), []byte(src), 0644); err != nil {
		t.Fatal(err)
	}

	result, diags := collectPackageTypes(dir)
	if len(diags) != 0 {
		t.Fatalf("diags = %d, want 0: %v", len(diags), diags)
	}
	if _, ok := result["FooRequest"]; !ok {
		t.Errorf("FooRequest not collected; keys=%v", keysOf(result))
	}
	if _, ok := result["FooResponse"]; !ok {
		t.Errorf("FooResponse not collected; keys=%v", keysOf(result))
	}
}

// TestCollectPackageTypesMissingDir — a non-existent path must be SILENT-OK (zero diagnostics).
func TestCollectPackageTypesMissingDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "does-not-exist")
	result, diags := collectPackageTypes(dir)
	if len(diags) != 0 {
		t.Fatalf("missing dir should be SILENT-OK, got diags=%v", diags)
	}
	if len(result) != 0 {
		t.Errorf("result should be empty, got keys=%v", keysOf(result))
	}
}

// TestCollectPackageTypesReadDirError — passing a file path where a directory is expected
// must report the ReadDir error as exactly one Diagnostic.
func TestCollectPackageTypesReadDirError(t *testing.T) {
	base := t.TempDir()
	filePath := filepath.Join(base, "not_a_dir.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	result, diags := collectPackageTypes(filePath)
	if len(diags) != 1 {
		t.Fatalf("diags = %d, want 1: %v", len(diags), diags)
	}
	d := diags[0]
	if d.File != filePath {
		t.Errorf("diag.File = %q, want %q", d.File, filePath)
	}
	if d.Phase != "parse" {
		t.Errorf("diag.Phase = %q, want parse", d.Phase)
	}
	if d.Level != "ERROR" {
		t.Errorf("diag.Level = %q, want ERROR", d.Level)
	}
	if !strings.Contains(d.Message, "cannot read funcspec type dir") {
		t.Errorf("diag.Message = %q, want ReadDir error message", d.Message)
	}
	if len(result) != 0 {
		t.Errorf("result should be empty on ReadDir error, got %v", keysOf(result))
	}
}

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

// TestFillMissingFieldsCacheDedup — when multiple specs share the same directory,
// a Diagnostic must be appended only once (seenDir deduplication).
func TestFillMissingFieldsCacheDedup(t *testing.T) {
	dir := t.TempDir()
	// A single file with a syntax error causes collectPackageTypes to emit one diag each call.
	bad := `package sample

!!!invalid!!!
`
	if err := os.WriteFile(filepath.Join(dir, "bad.go"), []byte(bad), 0644); err != nil {
		t.Fatal(err)
	}

	// Create two specs that reference the same dir.
	specs := []FuncSpec{
		{Name: "funcA", Package: "sample"},
		{Name: "funcB", Package: "sample"},
	}
	specDirs := []string{dir, dir}

	diags := fillMissingFields(specs, specDirs)
	if len(diags) != 1 {
		t.Fatalf("diags = %d, want 1 (dedup by dir): %v", len(diags), diags)
	}
}

func keysOf(m map[string][]Field) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
