//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what 단일 AST 필드에서 이름, 타입, JSON 태그를 추출하여 Field 목록을 반환한다
package funcspec

import (
	"go/ast"
)

func buildFields(f *ast.Field) []Field {
	typeName := exprToString(f.Type)
	jsonName := extractJSONTag(f)
	var fields []Field
	for _, name := range f.Names {
		fields = append(fields, Field{Name: name.Name, Type: typeName, JSONName: jsonName})
	}
	return fields
}
