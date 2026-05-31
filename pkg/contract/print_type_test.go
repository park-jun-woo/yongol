//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestPrintType — go/printer 로 AST 타입 표현식을 소스 문자열로 렌더 검증
package contract

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestPrintType(t *testing.T) {
	cases := []struct {
		name string
		expr string
		want string
	}{
		{"ident", "int", "int"},
		{"qualified", "context.Context", "context.Context"},
		{"pointer", "*Foo", "*Foo"},
		{"slice", "[]byte", "[]byte"},
		{"map", "map[string]int", "map[string]int"},
	}
	fset := token.NewFileSet()
	for _, c := range cases {
		expr, err := parser.ParseExpr(c.expr)
		if err != nil {
			t.Fatalf("%s: parse error: %v", c.name, err)
		}
		if got := printType(fset, expr); got != c.want {
			t.Errorf("%s: got %q want %q", c.name, got, c.want)
		}
	}
}
