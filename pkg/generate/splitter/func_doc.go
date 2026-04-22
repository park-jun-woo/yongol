//ff:func feature=gen-splitter type=util control=sequence
//ff:what funcDoc — FuncDecl 의 doc comment 텍스트 (없으면 "")
package splitter

import "go/ast"

// funcDoc returns the doc comment text attached to a FuncDecl. The
// returned string is the raw comment body (newline-joined) used later
// by summariseDoc for //ff:what extraction.
func funcDoc(d *ast.FuncDecl) string {
	if d.Doc == nil {
		return ""
	}
	return d.Doc.Text()
}
