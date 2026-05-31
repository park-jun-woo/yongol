//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagFloat64NumberWithFormatFloatReturnsFloatSpecificDiagnostic — float64 number with format float returns float-specific diagnostic 서브테스트
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagFloat64NumberWithFormatFloatReturnsFloatSpecificDiagnostic(t *testing.T) {

	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"number"}, Format: "float"}}
	diag, ok := xdo77PropDiag(xdo77FS(), "User", "users", "balance", ref, xdo77Cols())
	if !ok {
		t.Fatal("expected true for number/float on float64 column")
	}
	if !strings.Contains(diag.Message, "format: double") {
		t.Errorf("Message should mention format: double: %s", diag.Message)
	}

}
