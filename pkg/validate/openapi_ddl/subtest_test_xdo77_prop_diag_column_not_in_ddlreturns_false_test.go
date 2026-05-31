//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXdo77PropDiagColumnNotInDDLReturnsFalse — column not in DDL returns false 서브테스트
package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestXdo77PropDiagColumnNotInDDLReturnsFalse(t *testing.T) {

	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	_, ok := xdo77PropDiag(xdo77FS(), "User", "users", "ghost", ref, xdo77Cols())
	if ok {
		t.Error("expected false for missing column")
	}

}
