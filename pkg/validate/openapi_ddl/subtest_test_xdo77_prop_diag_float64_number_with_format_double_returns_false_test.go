//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagFloat64NumberWithFormatDoubleReturnsFalse — float64 number with format double returns false 서브테스트
package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagFloat64NumberWithFormatDoubleReturnsFalse(t *testing.T) {

	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"number"}, Format: "double"}}
	_, ok := xdo77PropDiag(xdo77FS(), "User", "users", "balance", ref, xdo77Cols())
	if ok {
		t.Error("expected false for number/double on float64 column")
	}

}
