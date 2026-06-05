//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestClassifyResponseShapes — classifyResponseShapes 단위 테스트 (embedded struct vs schema alias 분류)

package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClassifyResponseShapes(t *testing.T) {
	dir := t.TempDir()
	src := `package api

type GetWidget404JSONResponse ErrorResponse

type CreateWidget409JSONResponse struct{ ErrorJSONResponse }

type ErrorJSONResponse ErrorResponse
`
	if err := os.WriteFile(filepath.Join(dir, "responses.gen.go"), []byte(src), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	shapes := classifyResponseShapes(dir)

	alias, ok := shapes["GetWidget404JSONResponse"]
	if !ok {
		t.Fatalf("GetWidget404JSONResponse not classified")
	}
	if alias.Kind != shapeAlias {
		t.Errorf("GetWidget404JSONResponse kind = %v, want alias", alias.Kind)
	}

	emb, ok := shapes["CreateWidget409JSONResponse"]
	if !ok {
		t.Fatalf("CreateWidget409JSONResponse not classified")
	}
	if emb.Kind != shapeEmbedded {
		t.Errorf("CreateWidget409JSONResponse kind = %v, want embedded", emb.Kind)
	}
	if emb.EmbeddedType != "ErrorJSONResponse" {
		t.Errorf("CreateWidget409JSONResponse embedded = %q, want ErrorJSONResponse", emb.EmbeddedType)
	}
}
