//ff:func feature=contract type=util control=sequence
//ff:what renderRecv — SelectorExpr 의 수신자 쪽을 Ident 인 경우 이름, 그 외는 printer 로 문자열화

package contract

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

// renderRecv returns a printable string for the receiver side of a
// SelectorExpr. Simple Ident receivers return their name directly;
// other expressions are routed through go/printer so callers still
// see something useful for nested selectors. Errors from the printer
// collapse to the empty string so the caller can skip that entry.
func renderRecv(fset *token.FileSet, x ast.Expr) string {
	if ident, ok := x.(*ast.Ident); ok {
		return ident.Name
	}
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		return ""
	}
	return buf.String()
}
