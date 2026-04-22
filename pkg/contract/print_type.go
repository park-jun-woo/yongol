//ff:func feature=contract type=util control=sequence
//ff:what printType — Go AST 표현식을 go/printer 로 소스 문자열화

package contract

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

// printType renders a Go AST expression back to its canonical source
// string using go/printer. The output is used only for comparison,
// so the fallback on printer error returns an empty string rather
// than propagating it — an unprintable type is still a signal to the
// caller that something is off, but it does not abort the scan.
func printType(fset *token.FileSet, expr ast.Expr) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, expr); err != nil {
		return ""
	}
	return buf.String()
}
