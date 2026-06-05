//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestClassifyResponseShapesSkips — classifyResponseShapes 단위 테스트 (embedded struct vs schema alias 분류)

package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClassifyResponseShapesSkips(t *testing.T) {
	dir := t.TempDir()
	// sub-directory: skipped (e.IsDir()).
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	// non-.go file: skipped by suffix check.
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("hello"), 0o644); err != nil {
		t.Fatalf("write txt: %v", err)
	}
	// unparseable Go: parser error -> continue.
	if err := os.WriteFile(filepath.Join(dir, "broken.go"), []byte("package api\nthis is not go"), 0o644); err != nil {
		t.Fatalf("write broken: %v", err)
	}
	// valid Go with a non-type decl, a non-JSONResponse type, and a
	// JSONResponse type that classifyTypeSpec rejects (multi-field struct).
	src := `package api

var X = 1

type NotAResponse ErrorResponse

type WeirdJSONResponse struct {
	A int
	B int
}

type OkJSONResponse ErrorResponse
`
	if err := os.WriteFile(filepath.Join(dir, "mixed.go"), []byte(src), 0o644); err != nil {
		t.Fatalf("write mixed: %v", err)
	}

	shapes := classifyResponseShapes(dir)
	if _, ok := shapes["NotAResponse"]; ok {
		t.Error("NotAResponse should not be classified")
	}
	if _, ok := shapes["WeirdJSONResponse"]; ok {
		t.Error("multi-field struct should be rejected by classifyTypeSpec")
	}
	if s, ok := shapes["OkJSONResponse"]; !ok || s.Kind != shapeAlias {
		t.Errorf("OkJSONResponse = %v, ok=%v, want alias", s, ok)
	}
}

// TestClassifyTypeSpecDefault covers classifyTypeSpec's non-struct, non-ident
// branch (e.g. a map type) which returns ok=false.
