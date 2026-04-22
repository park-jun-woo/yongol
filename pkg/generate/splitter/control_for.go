//ff:func feature=gen-splitter type=util control=sequence
//ff:what controlFor — FuncDecl 이면 detectControl 호출, 그 외 선언은 sequence 반환
package splitter

import "go/ast"

// controlFor returns the (control, dimension) annotation tokens for the
// primary decl of a splitUnit. Only FuncDecl bodies are inspected; type
// and const/var decls always report "sequence" (they have no executable
// body) — callers may ignore the result for //ff:type files.
func controlFor(decl ast.Decl) (control, dimension string) {
	if fn, ok := decl.(*ast.FuncDecl); ok && fn.Body != nil {
		return detectControl(fn.Body)
	}
	return "sequence", ""
}
