//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what parseGoFile — preserved 파일 AST 파싱 (token.FileSet 동반 반환) helper

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// parseGoFile reads path and parses it with go/parser. It returns the
// parsed *ast.File together with its FileSet so callers can convert
// positions into human-readable line numbers.
//
// Parse errors short-circuit to (nil, nil, err); callers typically
// treat a parse failure as "out of scope" rather than a rule ERROR —
// PRV-01/02 already surface real syntax problems upstream.
func parseGoFile(path string) (*token.FileSet, *ast.File, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	return fset, file, nil
}
