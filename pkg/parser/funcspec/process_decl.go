//ff:func feature=funcspec type=parser control=selection
//ff:what AST 선언에서 Request/Response 구조체와 함수 본문을 추출한다
package funcspec

import (
	"go/ast"
	"go/token"
)

func processDecl(decl ast.Decl, fset *token.FileSet, spec *FuncSpec, expectedRequest, expectedResponse string) {
	switch d := decl.(type) {
	case *ast.GenDecl:
		if d.Tok == token.TYPE {
			processTypeSpecs(d, spec, expectedRequest, expectedResponse)
		}
	case *ast.FuncDecl:
		funcName := ucFirst(spec.Name)
		if d.Name.Name == funcName && d.Body != nil {
			spec.HasBody = !isStubBody(fset, d.Body)
		}
	}
}
