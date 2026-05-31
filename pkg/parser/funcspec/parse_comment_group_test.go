//ff:func feature=funcspec type=test control=sequence
//ff:what applyAnnotation / parseCommentGroup / countFuncAnnotations — @func/@error/@description
package funcspec

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestParseCommentGroup(t *testing.T) {
	src := `package p

// @func hashPassword
// @error 401
// @description does a thing
func hashPassword() {}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var s FuncSpec
	parseCommentGroup(f.Comments[0], fset, &s)
	if s.Name != "hashPassword" || s.ErrStatus != 401 || s.Description != "does a thing" {
		t.Errorf("spec = %+v", s)
	}
}
