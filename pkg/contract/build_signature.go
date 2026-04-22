//ff:func feature=contract type=util control=sequence
//ff:what buildSignature — FuncDecl 을 FuncSignature 구조체로 변환 (파라미터·반환·HasErr)

package contract

import (
	"go/ast"
	"go/token"
)

// buildSignature converts a FuncDecl to FuncSignature by delegating
// parameter and return-type expansion to their own helpers. HasErr
// is set when the last rendered return type is exactly "error".
func buildSignature(fset *token.FileSet, fd *ast.FuncDecl) FuncSignature {
	sig := FuncSignature{Name: fd.Name.Name}
	sig.Params = expandFieldList(fset, fd.Type.Params, true)
	if returnsList := expandFieldList(fset, fd.Type.Results, false); len(returnsList) > 0 {
		sig.Returns = make([]string, 0, len(returnsList))
		for _, p := range returnsList {
			sig.Returns = append(sig.Returns, p.Type)
		}
	}
	if n := len(sig.Returns); n > 0 && sig.Returns[n-1] == "error" {
		sig.HasErr = true
	}
	return sig
}
