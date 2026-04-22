//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what collectImportsForDecl — 선언이 실제 사용하는 package 식별자를 찾아 관련 ImportSpec 목록만 반환
package splitter

import "go/ast"

// collectImportsForDecl walks decls and returns only those ImportSpec
// entries whose exported package name is referenced via a SelectorExpr
// (pkg.Name). Dot imports and blank imports (_ / .) are always included
// so side-effects survive the split. The matching uses the import alias
// when present, otherwise the last segment of the import path — the same
// resolution Go uses for selectors.
func collectImportsForDecl(decls []ast.Decl, allImports []*ast.ImportSpec) []*ast.ImportSpec {
	used := map[string]bool{}
	for _, d := range decls {
		gatherSelectorNames(d, used)
	}
	var out []*ast.ImportSpec
	for _, imp := range allImports {
		if keepImport(imp, used) {
			out = append(out, imp)
		}
	}
	return out
}
