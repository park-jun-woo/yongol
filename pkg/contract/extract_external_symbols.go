//ff:func feature=contract type=util control=iteration dimension=1
//ff:what ExtractExternalSymbols — 파일의 첫 func body 를 AST walk 해 외부 심볼을 분류 수집

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// ExtractExternalSymbols parses filePath and walks the body of its
// first non-init FuncDecl, classifying every SelectorExpr it sees
// into one of three buckets (sqlc queries, package calls, struct
// field access). Files without a qualifying func return the zero
// ExternalSymbols (nil error).
func ExtractExternalSymbols(filePath string) (ExternalSymbols, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return ExternalSymbols{}, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, src, parser.ParseComments)
	if err != nil {
		return ExternalSymbols{}, err
	}
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Name.Name == "init" || fd.Body == nil {
			continue
		}
		return walkForSymbols(fset, fd.Body), nil
	}
	return ExternalSymbols{}, nil
}
