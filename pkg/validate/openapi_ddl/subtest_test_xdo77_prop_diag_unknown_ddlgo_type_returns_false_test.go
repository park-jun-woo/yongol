//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagUnknownDDLGoTypeReturnsFalse — unknown DDL Go type returns false 서브테스트
package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagUnknownDDLGoTypeReturnsFalse(t *testing.T) {

	unknownCols := map[string]string{"data": "unknownType"}
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	_, ok := xdo77PropDiag(xdo77FS(), "User", "users", "data", ref, unknownCols)
	if ok {
		t.Error("expected false for unknown DDL Go type")
	}

}
