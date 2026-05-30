//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestUcFirst(t *testing.T) {
	tests := []struct{ in, want string }{
		{"hashPassword", "HashPassword"},
		{"user_id", "UserID"}, // strcase ToGoPascal applies initialism handling
		{"id", "ID"},
	}
	for _, tt := range tests {
		if got := ucFirst(tt.in); got != tt.want {
			t.Errorf("ucFirst(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestAppendN(t *testing.T) {
	got := appendN(nil, "int", 3)
	want := []string{"int", "int", "int"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("appendN = %v, want %v", got, want)
	}
	// n=0 → unchanged.
	base := []string{"a"}
	if got := appendN(base, "x", 0); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("appendN n=0 = %v", got)
	}
}

// parseExpr is a small helper that parses a Go expression into an ast.Expr.
func parseExpr(t *testing.T, src string) ast.Expr {
	t.Helper()
	e, err := parser.ParseExpr(src)
	if err != nil {
		t.Fatalf("parse %q: %v", src, err)
	}
	return e
}

func TestIsZeroExpr(t *testing.T) {
	tests := []struct {
		src  string
		want bool
	}{
		{"nil", true},
		{"false", true},
		{"0", true},
		{"0.0", true},
		{`""`, true},
		{"Resp{}", true},
		{"42", false},
		{`"hello"`, false},
		{"true", false},
		{"Resp{Status: 1}", false},
		{"someVar", false},
	}
	for _, tt := range tests {
		if got := isZeroExpr(parseExpr(t, tt.src)); got != tt.want {
			t.Errorf("isZeroExpr(%q) = %v, want %v", tt.src, got, tt.want)
		}
	}
}

func TestAllZeroResults(t *testing.T) {
	if !allZeroResults([]ast.Expr{parseExpr(t, "nil"), parseExpr(t, "Resp{}")}) {
		t.Error("all-zero results should be true")
	}
	if allZeroResults([]ast.Expr{parseExpr(t, "nil"), parseExpr(t, "42")}) {
		t.Error("mixed results should be false")
	}
	// empty list → true.
	if !allZeroResults(nil) {
		t.Error("empty results should be true")
	}
}

// firstStmt parses a function body and returns its first statement.
func parseBody(t *testing.T, body string) *ast.BlockStmt {
	t.Helper()
	src := "package p\nfunc f() (Resp, error) {\n" + body + "\n}\n"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse body: %v", err)
	}
	return file.Decls[0].(*ast.FuncDecl).Body
}

func TestIsPanicCall(t *testing.T) {
	body := parseBody(t, `panic("TODO")`)
	if !isPanicCall(body.List[0]) {
		t.Error("expected panic call detected")
	}
	body2 := parseBody(t, `return Resp{}, nil`)
	if isPanicCall(body2.List[0]) {
		t.Error("return is not a panic call")
	}
}

func TestIsStubBody(t *testing.T) {
	fset := token.NewFileSet()
	tests := []struct {
		name string
		body string
		want bool
	}{
		{"empty", ``, true},
		{"panic", `panic("TODO")`, true},
		{"zero return", `return Resp{}, nil`, true},
		{"meaningful return", `return Resp{Status: 1}, nil`, false},
		{"multi stmt", "x := 1\nreturn Resp{}, nil", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStubBody(fset, parseBody(t, tt.body)); got != tt.want {
				t.Errorf("isStubBody(%q) = %v, want %v", tt.body, got, tt.want)
			}
		})
	}
}

func TestExtractJSONTag(t *testing.T) {
	field := func(tag string) *ast.Field {
		var f ast.Field
		if tag != "" {
			f.Tag = &ast.BasicLit{Kind: token.STRING, Value: "`" + tag + "`"}
		}
		return &f
	}
	if got := extractJSONTag(field(`json:"user_id"`)); got != "user_id" {
		t.Errorf("got %q, want user_id", got)
	}
	if got := extractJSONTag(field(`json:"name,omitempty"`)); got != "name" {
		t.Errorf("got %q, want name (option stripped)", got)
	}
	if got := extractJSONTag(field(`json:"-"`)); got != "" {
		t.Errorf("json:- → %q, want empty", got)
	}
	if got := extractJSONTag(field("")); got != "" {
		t.Errorf("no tag → %q, want empty", got)
	}
	if got := extractJSONTag(field(`xml:"x"`)); got != "" {
		t.Errorf("no json tag → %q, want empty", got)
	}
}

func TestExprToString(t *testing.T) {
	tests := []struct{ src, want string }{
		{"int", "int"},
		{"pkg.Type", "pkg.Type"},
		{"*User", "*User"},
		{"[]string", "[]string"},
		{"map[string]int", "map[string]int"},
		{"[]*pkg.Type", "[]*pkg.Type"},
	}
	for _, tt := range tests {
		if got := exprToString(parseExpr(t, tt.src)); got != tt.want {
			t.Errorf("exprToString(%q) = %q, want %q", tt.src, got, tt.want)
		}
	}
	// Unknown expr falls back to interface{}.
	if got := exprToString(parseExpr(t, "func(){}")); got != "interface{}" {
		t.Errorf("func type → %q, want interface{}", got)
	}
}

func TestExtractGoParseErrorLine(t *testing.T) {
	// Parsing invalid Go produces a scanner.ErrorList with a line.
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "bad.go", "package p\nfunc {\n", 0)
	if err == nil {
		t.Fatal("expected a parse error")
	}
	if line := extractGoParseErrorLine(err); line <= 0 {
		t.Errorf("expected a positive error line, got %d", line)
	}
	// Non-scanner error → 0.
	if line := extractGoParseErrorLine(errPlain{}); line != 0 {
		t.Errorf("non-scanner error → %d, want 0", line)
	}
}

type errPlain struct{}

func (errPlain) Error() string { return "x" }
