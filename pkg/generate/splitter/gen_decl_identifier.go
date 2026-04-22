//ff:func feature=gen-splitter type=util control=sequence
//ff:what genDeclIdentifier — GenDecl 식별자 (단일 spec 이름 또는 token 문자열)
package splitter

import "go/ast"

// genDeclIdentifier returns the representative name for a GenDecl.
// Single-spec decls expose their declared name; everything else falls
// back to the GenDecl token ("const", "var", "type") so //ff:what is
// never empty.
func genDeclIdentifier(d *ast.GenDecl) string {
	if n := primaryTypeName(d); n != "" {
		return n
	}
	return d.Tok.String()
}
