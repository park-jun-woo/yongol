//ff:func feature=gen-splitter type=util control=sequence
//ff:what renderImportSpec — ImportSpec 하나를 "alias \"path\"" 또는 "\"path\"" 문자열로 변환
package splitter

import "go/ast"

// renderImportSpec formats a single ImportSpec into the exact textual
// form Go source uses: "alias \"path\"" when an alias (including _ or
// . alias) is present, otherwise just "\"path\"".
func renderImportSpec(imp *ast.ImportSpec) string {
	if imp.Name == nil {
		return imp.Path.Value
	}
	return imp.Name.Name + " " + imp.Path.Value
}
