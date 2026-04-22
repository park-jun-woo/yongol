//ff:func feature=gen-splitter type=util control=sequence
//ff:what sqlcSuffix — sqlc 산출 decl 에 대한 ".sql.go" vs ".model.go" 선택
package splitter

import "go/ast"

// sqlcSuffix returns ".model.go" when the decl is a type declaration
// originating from sqlc's models.go, else ".sql.go". Methods on row
// types, query funcs and param structs all land on .sql.go so sqlc's
// runtime convention (one .sql.go per source query file, split per
// method by us) stays visible.
func sqlcSuffix(isModelsFile bool, decl ast.Decl) string {
	if !isModelsFile {
		return ".sql.go"
	}
	gd, ok := decl.(*ast.GenDecl)
	if !ok || gd.Tok.String() != "type" {
		return ".sql.go"
	}
	return ".model.go"
}
