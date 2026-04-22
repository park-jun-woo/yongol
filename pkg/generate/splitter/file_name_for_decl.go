//ff:func feature=gen-splitter type=util control=selection
//ff:what fileNameForDecl — 선언 종류에 맞춰 tool 관례의 파일명(snake_case) 생성
package splitter

import "go/ast"

// fileNameForDecl maps a single ast.Decl to the target split file name
// (without directory). It respects each tool's suffix convention:
//
//	oapi-codegen  → *.gen.go     (all splits keep the .gen.go suffix)
//	sqlc          → *.sql.go     for funcs/methods and types outside models.go
//	sqlc          → *.model.go   for struct/alias types originating from models.go
//
// Callers pass isModelsFile=true when the source file is sqlc's models.go.
// Name normalisation is CamelCase → snake_case (see snake()).
func fileNameForDecl(decl ast.Decl, tool Tool, isModelsFile bool) string {
	suffix := suffixFor(tool, isModelsFile, decl)
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return funcFileName(d, suffix)
	case *ast.GenDecl:
		return genDeclFileName(d, suffix)
	}
	return "misc" + suffix
}
