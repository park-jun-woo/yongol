//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what AST 구조체에서 필드 목록을 추출한다
package funcspec

import (
	"go/ast"
)

// extractFields extracts field names and types from a struct.
func extractFields(st *ast.StructType) []Field {
	var fields []Field
	for _, f := range st.Fields.List {
		fields = append(fields, buildFields(f)...)
	}
	return fields
}
