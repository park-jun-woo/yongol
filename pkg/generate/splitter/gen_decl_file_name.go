//ff:func feature=gen-splitter type=util control=selection
//ff:what genDeclFileName — GenDecl(type/const/var) 의 파일명 계산
package splitter

import "go/ast"

// genDeclFileName returns the split file name for a GenDecl. Single-spec
// type/const/var decls use their declared name; multi-spec const or var
// blocks land in the shared consts / vars bucket files so they stay
// together (both are legal under filefunc F5).
func genDeclFileName(d *ast.GenDecl, suffix string) string {
	if name := primaryTypeName(d); name != "" {
		return snake(name) + suffix
	}
	switch d.Tok.String() {
	case "const":
		return "consts" + suffix
	case "var":
		return "vars" + suffix
	}
	return "misc" + suffix
}
