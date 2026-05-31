//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagTypeMismatchReturnsDiagnostic — type mismatch returns diagnostic 서브테스트
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagTypeMismatchReturnsDiagnostic(t *testing.T) {

	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	diag, ok := xdo77PropDiag(xdo77FS(), "User", "users", "id", ref, xdo77Cols())
	if !ok {
		t.Fatal("expected true for mismatch")
	}
	if !strings.Contains(diag.Message, "XDO-77") {
		t.Errorf("Message missing XDO-77: %s", diag.Message)
	}

}
