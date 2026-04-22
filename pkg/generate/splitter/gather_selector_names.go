//ff:func feature=gen-splitter type=util control=sequence
//ff:what gatherSelectorNames — ast.Inspect 로 decl 하위에서 pkg.Name 형태 selector 식별자 수집
package splitter

import "go/ast"

// gatherSelectorNames walks decl and records into used every identifier
// that appears as the package qualifier of a SelectorExpr (the "pkg" in
// pkg.Name). Identifiers resolved to a local object are skipped so
// local names shadowing a package do not leak into the import set.
func gatherSelectorNames(decl ast.Decl, used map[string]bool) {
	ast.Inspect(decl, func(n ast.Node) bool {
		id := selectorPkgIdent(n)
		if id != nil && id.Obj == nil {
			used[id.Name] = true
		}
		return true
	})
}
