//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestErrorResponseLiteralEmbedded — errorResponseLiteral 단위 테스트 (alias 폴백 vs embedded 래퍼 리터럴)

package ssac

import "testing"

func TestErrorResponseLiteralEmbedded(t *testing.T) {
	g := &methodGen{
		FuncName: "CreateWidget",
		ResponseShapes: map[string]respShape{
			"CreateWidget409JSONResponse": {Kind: shapeEmbedded, EmbeddedType: "ErrorJSONResponse"},
		},
	}
	got := g.errorResponseLiteral(409, "Conflict", "conflict")
	want := `api.CreateWidget409JSONResponse{ErrorJSONResponse: api.ErrorJSONResponse{Error: "Conflict", Code: "conflict"}}`
	if got != want {
		t.Errorf("errorResponseLiteral = %q, want %q", got, want)
	}
}
