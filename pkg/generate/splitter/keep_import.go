//ff:func feature=gen-splitter type=util control=sequence
//ff:what keepImport — ImportSpec 을 분할 파일에 유지해야 하는지 판정 (blank/dot/사용된 식별자)
package splitter

import "go/ast"

// keepImport returns true when the ImportSpec must survive the split:
// blank (_) and dot (.) imports always stay (side-effect coupling),
// named imports stay only if used tracks their binding name.
func keepImport(imp *ast.ImportSpec, used map[string]bool) bool {
	if imp.Name != nil && (imp.Name.Name == "_" || imp.Name.Name == ".") {
		return true
	}
	return used[importName(imp)]
}
