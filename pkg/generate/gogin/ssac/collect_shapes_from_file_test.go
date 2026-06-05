//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectShapesFromFile — collectShapesFromFile 단위 테스트 (embedded/alias 적재, 비-JSONResponse·비-TYPE 선언 skip)

package ssac

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCollectShapesFromFile(t *testing.T) {
	src := `package api

func helper() {}

type GetWidget404JSONResponse ErrorResponse

type CreateWidget409JSONResponse struct{ ErrorJSONResponse }

type ErrorJSONResponse ErrorResponse

type NotAWrapper struct{ X int }
`
	file, err := parser.ParseFile(token.NewFileSet(), "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	shapes := map[string]respShape{}
	collectShapesFromFile(file, shapes)

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

	if _, ok := shapes["NotAWrapper"]; ok {
		t.Errorf("NotAWrapper should be skipped (no JSONResponse suffix)")
	}
}
