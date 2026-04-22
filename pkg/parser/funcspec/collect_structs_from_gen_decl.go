//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what 단일 GenDecl에서 구조체 spec을 추출하여 result 맵에 추가한다
package funcspec

import "go/ast"

func collectStructsFromGenDecl(gd *ast.GenDecl, result map[string][]Field) {
	for _, spec := range gd.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}
		result[ts.Name.Name] = extractFields(st)
	}
}
