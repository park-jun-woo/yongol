//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what applyAnnotation / parseCommentGroup / countFuncAnnotations — @func/@error/@description

package funcspec

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestApplyAnnotation(t *testing.T) {
	t.Run("func sets name and line", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@func hashPassword", 7, &s)
		if s.Name != "hashPassword" || s.Line != 7 {
			t.Errorf("spec = %+v", s)
		}
	})
	t.Run("error sets status", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@error 422", 1, &s)
		if s.ErrStatus != 422 {
			t.Errorf("ErrStatus = %d, want 422", s.ErrStatus)
		}
	})
	t.Run("error non-numeric ignored", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@error notanumber", 1, &s)
		if s.ErrStatus != 0 {
			t.Errorf("ErrStatus = %d, want 0", s.ErrStatus)
		}
	})
	t.Run("description", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@description hashes a password", 1, &s)
		if s.Description != "hashes a password" {
			t.Errorf("Description = %q", s.Description)
		}
	})
	t.Run("unknown ignored", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@unknown x", 1, &s)
		if s.Name != "" || s.Description != "" || s.ErrStatus != 0 {
			t.Errorf("spec should be empty: %+v", s)
		}
	})
}

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
