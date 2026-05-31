//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagFormatMismatchWithFormatInDisplay — format mismatch with format in display 서브테스트
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagFormatMismatchWithFormatInDisplay(t *testing.T) {

	// id is int64 which expects integer/int64, provide integer/int32
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int32"}}
	diag, ok := xdo77PropDiag(xdo77FS(), "User", "users", "id", ref, xdo77Cols())
	if !ok {
		t.Fatal("expected true for format mismatch")
	}
	if !strings.Contains(diag.Message, "int32") {
		t.Errorf("Message missing format: %s", diag.Message)
	}

}
