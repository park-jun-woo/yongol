//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what 단일 AST 파일에서 구조체 타입을 수집하여 result 맵에 추가한다
package funcspec

import (
	"go/ast"
	"go/token"
)

func collectStructsFromFile(f *ast.File, result map[string][]Field) {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		collectStructsFromGenDecl(gd, result)
	}
}
