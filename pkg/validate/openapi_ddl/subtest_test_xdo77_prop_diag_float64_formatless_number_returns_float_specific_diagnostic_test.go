//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagFloat64FormatlessNumberReturnsFloatSpecificDiagnostic — float64 formatless number returns float-specific diagnostic 서브테스트
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagFloat64FormatlessNumberReturnsFloatSpecificDiagnostic(t *testing.T) {

	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"number"}}}
	diag, ok := xdo77PropDiag(xdo77FS(), "User", "users", "balance", ref, xdo77Cols())
	if !ok {
		t.Fatal("expected true for formatless number on float64 column")
	}
	if !strings.Contains(diag.Message, "format: double") {
		t.Errorf("Message should mention format: double: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "float64") {
		t.Errorf("Message should mention float64: %s", diag.Message)
	}

}
