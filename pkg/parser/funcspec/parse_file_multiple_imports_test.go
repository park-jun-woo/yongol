//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what ParseFile 복수 import 파싱 테스트 — grouped import 블록에서 3개 경로 수집 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileMultipleImports(t *testing.T) {
	dir := t.TempDir()

	src := `package bad

import (
	"database/sql"
	"fmt"
	"net/http"
)

// @func badFunc
// @description does bad things

type BadFuncRequest struct{}
type BadFuncResponse struct{}
func BadFunc(req BadFuncRequest) (BadFuncResponse, error) {
	return BadFuncResponse{}, nil
}
`
	path := filepath.Join(dir, "bad_func.go")
	os.WriteFile(path, []byte(src), 0644)

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	if len(spec.Imports) != 3 {
		t.Fatalf("Imports count = %d, want 3", len(spec.Imports))
	}

	expected := map[string]bool{"database/sql": true, "fmt": true, "net/http": true}
	for _, imp := range spec.Imports {
		if !expected[imp] {
			t.Errorf("unexpected import: %q", imp)
		}
	}
}
