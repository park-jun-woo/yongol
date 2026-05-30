//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestExtractQuoted(t *testing.T) {
	val, rest := extractQuoted(`"hello" world`)
	if val != "hello" || rest != "world" {
		t.Errorf("extractQuoted = (%q,%q), want (hello,world)", val, rest)
	}
	// No leading quote → ("", original-trimmed).
	v2, r2 := extractQuoted(`bare value`)
	if v2 != "" || r2 != "bare value" {
		t.Errorf("no-quote = (%q,%q)", v2, r2)
	}
	// Unterminated quote → ("", original).
	v3, r3 := extractQuoted(`"unterminated`)
	if v3 != "" || r3 != `"unterminated` {
		t.Errorf("unterminated = (%q,%q)", v3, r3)
	}
}

func TestIsLiteral(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"true", true},
		{"false", true},
		{"nil", true},
		{"42", true},
		{"-1", true},
		{"3.14", true},
		{"-0.5", true},
		{"", false},
		{"-", false},
		{"foo", false},
		{"1.2.3", false}, // second dot → not numeric
		{"course.ID", false},
	}
	for _, tt := range tests {
		if got := IsLiteral(tt.in); got != tt.want {
			t.Errorf("IsLiteral(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestSplitPackagePrefix(t *testing.T) {
	tests := []struct {
		in       string
		wantPkg  string
		wantRest string
	}{
		{"session.Session.Get", "session", "Session.Get"},
		{"Course.FindByID", "", "Course.FindByID"},   // 2-part, no pkg
		{"plain", "", "plain"},                       // no dot
		{"Pkg.Model.Method", "", "Pkg.Model.Method"}, // uppercase first → no pkg
	}
	for _, tt := range tests {
		pkg, rest := splitPackagePrefix(tt.in)
		if pkg != tt.wantPkg || rest != tt.wantRest {
			t.Errorf("splitPackagePrefix(%q) = (%q,%q), want (%q,%q)", tt.in, pkg, rest, tt.wantPkg, tt.wantRest)
		}
	}
}

func TestSplitTargetMessage(t *testing.T) {
	target, msg, rem := splitTargetMessage(`course "not found" 404`)
	if target != "course" || msg != "not found" || rem != "404" {
		t.Errorf("got (%q,%q,%q)", target, msg, rem)
	}
	// No quote → target only.
	t2, m2, r2 := splitTargetMessage(`bareTarget`)
	if t2 != "bareTarget" || m2 != "" || r2 != "" {
		t.Errorf("no-quote = (%q,%q,%q)", t2, m2, r2)
	}
}

func TestParseTwoQuoted(t *testing.T) {
	first, second, rem := parseTwoQuoted(`"a" "b" tail`)
	if first != "a" || second != "b" || rem != "tail" {
		t.Errorf("got (%q,%q,%q)", first, second, rem)
	}
	// Only one quoted → second empty.
	f2, s2, r2 := parseTwoQuoted(`"only"`)
	if f2 != "only" || s2 != "" || r2 != "" {
		t.Errorf("one-quoted = (%q,%q,%q)", f2, s2, r2)
	}
}

func TestHasNoPaginationComment(t *testing.T) {
	yes := []*ast.Comment{
		{Text: "// some doc"},
		{Text: "// @no-pagination"},
	}
	if !hasNoPaginationComment(yes) {
		t.Error("expected @no-pagination detected")
	}
	no := []*ast.Comment{{Text: "// @something-else"}}
	if hasNoPaginationComment(no) {
		t.Error("should not detect @no-pagination")
	}
	if hasNoPaginationComment(nil) {
		t.Error("nil comments → false")
	}
}

func TestExprToString(t *testing.T) {
	mustExpr := func(src string) ast.Expr {
		e, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatalf("parse %q: %v", src, err)
		}
		return e
	}
	if got := exprToString(mustExpr("foo")); got != "foo" {
		t.Errorf("ident = %q", got)
	}
	if got := exprToString(mustExpr("pkg.Type")); got != "pkg.Type" {
		t.Errorf("selector = %q", got)
	}
	if got := exprToString(mustExpr("a.b.c")); got != "a.b.c" {
		t.Errorf("nested selector = %q", got)
	}
	// Unsupported expr → "".
	if got := exprToString(mustExpr("[]int")); got != "" {
		t.Errorf("array expr → %q, want empty", got)
	}
}
