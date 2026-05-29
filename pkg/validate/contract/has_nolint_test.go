//ff:func feature=validate-contract type=test control=iteration dimension=2 topic=preserve-safety
//ff:what TestHasNolint — 특정 라인(또는 바로 위) nolint 주석 존재 여부 검증

package contract

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestHasNolint(t *testing.T) {
	src := "package p\n" + // line 1
		"\n" + // line 2
		"func f() { // nolint:prv-13\n" + // line 3
		"\t// nolint:panic\n" + // line 4
		"\tdoThing()\n" + // line 5
		"}\n" // line 6

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		line int
		rule string
		want bool
	}{
		{"same line", 3, "prv-13", true},
		{"line above", 5, "panic", true},
		{"wrong rule same line", 3, "panic", false},
		{"unrelated line", 6, "prv-13", false},
		{"line zero rejected", 0, "panic", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasNolint(fset, file, tt.line, tt.rule); got != tt.want {
				t.Fatalf("hasNolint(line=%d, %q) = %v, want %v", tt.line, tt.rule, got, tt.want)
			}
		})
	}

	if hasNolint(nil, file, 3, "prv-13") {
		t.Fatal("nil fset should return false")
	}
	if hasNolint(fset, nil, 3, "prv-13") {
		t.Fatal("nil file should return false")
	}
}
