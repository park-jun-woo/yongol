//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagFloat64WrongTypeReturnsGenericDiagnostic — float64 wrong type returns generic diagnostic 서브테스트
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagFloat64WrongTypeReturnsGenericDiagnostic(t *testing.T) {

	// type mismatch (string vs number) takes the generic path, not the float helper.
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	diag, ok := xdo77PropDiag(xdo77FS(), "User", "users", "balance", ref, xdo77Cols())
	if !ok {
		t.Fatal("expected true for string on float64 column")
	}
	if !strings.Contains(diag.Message, "XDO-77") {
		t.Errorf("Message missing XDO-77: %s", diag.Message)
	}

}
