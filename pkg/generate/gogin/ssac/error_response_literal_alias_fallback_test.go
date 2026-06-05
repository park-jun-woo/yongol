//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestErrorResponseLiteralAliasFallback — errorResponseLiteral 단위 테스트 (alias 폴백 vs embedded 래퍼 리터럴)

package ssac

import "testing"

func TestErrorResponseLiteralAliasFallback(t *testing.T) {
	// No ResponseShapes entry → alias form (prior behaviour preserved).
	g := &methodGen{FuncName: "GetWidget"}
	got := g.errorResponseLiteral(404, "Not found", "not_found")
	want := `api.GetWidget404JSONResponse{Error: "Not found", Code: "not_found"}`
	if got != want {
		t.Errorf("errorResponseLiteral = %q, want %q", got, want)
	}
}
