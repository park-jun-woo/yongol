//ff:func feature=funcspec type=test control=sequence
//ff:what applyAnnotation / parseCommentGroup / countFuncAnnotations — @func/@error/@description
package funcspec

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCountFuncAnnotations(t *testing.T) {
	src := `package p

// @func one
// not annotation

// @func two
func f() {}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if n := countFuncAnnotations(f.Comments); n != 2 {
		t.Errorf("countFuncAnnotations = %d, want 2", n)
	}
}
