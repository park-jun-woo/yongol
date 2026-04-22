//ff:func feature=gen-splitter type=util control=selection
//ff:what docOf — FuncDecl / GenDecl 의 doc comment 텍스트(없으면 "") 반환
package splitter

import "go/ast"

// docOf returns the raw doc comment text associated with a decl.
// summariseDoc later extracts the first non-blank line for //ff:what.
func docOf(decl ast.Decl) string {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return funcDoc(d)
	case *ast.GenDecl:
		return genDeclDoc(d)
	}
	return ""
}
