//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what collectFileImports — 파일 import 목록에서 패키지 이름 집합 추출

package contract

import (
	"go/parser"
	"go/token"
)

// collectFileImports parses filePath and returns the set of package
// names it imports (the right-most segment of each import path, or the
// alias when one is present). PRV-02 uses this to distinguish
// `pkg.Symbol` package accesses (which are NOT DDL columns) from
// `recv.Field` struct-field accesses.
//
// Errors parsing the file yield an empty set — PRV-02 simply loses
// its filter and treats everything literally, which is safe.
func collectFileImports(filePath string) map[string]bool {
	pkgs := map[string]bool{}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ImportsOnly)
	if err != nil {
		return pkgs
	}
	for _, imp := range f.Imports {
		pkgs[resolveImportName(imp.Name, imp.Path.Value)] = true
	}
	return pkgs
}
