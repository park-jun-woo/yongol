//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestErrorResponseLiteralAliasKindExplicit — errorResponseLiteral 단위 테스트 (alias 폴백 vs embedded 래퍼 리터럴)

package ssac

import "testing"

func TestErrorResponseLiteralAliasKindExplicit(t *testing.T) {
	// An explicit alias classification also emits the flat form.
	g := &methodGen{
		FuncName: "GetWidget",
		ResponseShapes: map[string]respShape{
			"GetWidget404JSONResponse": {Kind: shapeAlias},
		},
	}
	got := g.errorResponseLiteral(404, "Not found", "not_found")
	want := `api.GetWidget404JSONResponse{Error: "Not found", Code: "not_found"}`
	if got != want {
		t.Errorf("errorResponseLiteral = %q, want %q", got, want)
	}
}
