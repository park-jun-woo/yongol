//ff:func feature=funcspec type=test control=sequence
//ff:what collectPackageTypes — 정상 Go 파일에서 struct 2개 수집 (진단 0)

package funcspec

import (
	"os"
	"path/filepath"
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
