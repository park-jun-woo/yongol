//ff:func feature=ssac-parse type=test-helper control=iteration dimension=1
//ff:what exerciseGenDeclSpecs — GenDecl 의 각 spec 에 extractStructInfo 직접 호출 헬퍼
package ssac

import "go/ast"

// exerciseGenDeclSpecs calls extractStructInfo on each spec of a GenDecl.
func exerciseGenDeclSpecs(decl ast.Decl) {
	gd, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	for _, spec := range gd.Specs {
		extractStructInfo(spec)
	}
}
