//ff:func feature=gen-splitter type=util control=selection
//ff:what declIdentifier — FuncDecl/GenDecl 에서 대표 식별자 이름 추출
package splitter

import "go/ast"

// declIdentifier returns the primary name for a decl, used as the
// fallback //ff:what summary when the source carries no doc comment.
// For methods the receiver is prefixed (e.g. "Queries.UserFindByEmail").
func declIdentifier(decl ast.Decl) string {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return funcIdentifier(d)
	case *ast.GenDecl:
		return genDeclIdentifier(d)
	}
	return ""
}
