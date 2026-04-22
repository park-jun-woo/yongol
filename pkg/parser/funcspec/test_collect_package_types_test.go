//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what collectPackageTypes / fillMissingFields 의 Diagnostic 전파 동작을 검증한다

package funcspec

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCollectPackageTypesNormal — 정상 Go 파일은 struct 를 수집하고 diag 이 없어야 한다.
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

// TestCollectPackageTypesMissingDir — 존재하지 않는 경로는 SILENT-OK (diag 0).
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

// TestCollectPackageTypesReadDirError — 파일을 디렉토리 경로로 넘기면 ReadDir 에러
// 가 Diagnostic 1 건으로 보고되어야 한다.
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

// TestCollectPackageTypesSyntaxErrorPartial — 문법 오류 파일은 diag 1 건으로 보고되고
// 같은 디렉토리의 다른 정상 파일은 여전히 수집되어야 한다 (partial success).
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
	// partial success: ok.go 의 OkRequest 는 여전히 수집되어야 한다.
	if _, found := result["OkRequest"]; !found {
		t.Errorf("OkRequest should still be collected despite sibling syntax error; keys=%v", keysOf(result))
	}
}

// TestFillMissingFieldsCacheDedup — 같은 디렉토리를 여러 spec 이 공유하면
// Diagnostic 은 1 회만 append 되어야 한다 (seenDir dedup).
func TestFillMissingFieldsCacheDedup(t *testing.T) {
	dir := t.TempDir()
	// 문법 오류 파일 하나만 두면 collectPackageTypes 가 매번 diag 1 건을 낸다.
	bad := `package sample

!!!invalid!!!
`
	if err := os.WriteFile(filepath.Join(dir, "bad.go"), []byte(bad), 0644); err != nil {
		t.Fatal(err)
	}

	// 같은 dir 을 참조하는 spec 을 2 개 만든다.
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
