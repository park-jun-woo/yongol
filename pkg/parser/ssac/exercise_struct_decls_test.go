//ff:func feature=ssac-parse type=test-helper control=iteration dimension=1
//ff:what exerciseStructDecls — 파일 선언에서 collectStructsFromDecl + extractStructInfo 직접 호출 헬퍼
package ssac

import "go/ast"

// exerciseStructDecls calls collectStructsFromDecl on every decl and
// extractStructInfo on every type spec, for coverage attribution.
func exerciseStructDecls(f *ast.File) {
	for _, decl := range f.Decls {
		collectStructsFromDecl(decl)
		exerciseGenDeclSpecs(decl)
	}
}
