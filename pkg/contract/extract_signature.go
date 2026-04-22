//ff:func feature=contract type=util control=iteration dimension=1
//ff:what ExtractSignature — 파일을 파싱해 첫 func(init 제외)의 FuncSignature 를 반환

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// ExtractSignature parses filePath and returns the FuncSignature of
// the first top-level FuncDecl whose name is not "init". Files with
// no qualifying func yield the zero FuncSignature (nil error) so
// callers can treat them as "no signature to compare".
//
// Parameter and return type expressions are rendered back to source
// via go/printer so the string form matches the author's original
// spelling (including qualified identifiers such as `pkg.Type` and
// generic instantiations). HasErr is set when the last return type
// prints as exactly "error".
func ExtractSignature(filePath string) (FuncSignature, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return FuncSignature{}, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, src, parser.ParseComments)
	if err != nil {
		return FuncSignature{}, err
	}
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Name.Name == "init" {
			continue
		}
		return buildSignature(fset, fd), nil
	}
	return FuncSignature{}, nil
}
