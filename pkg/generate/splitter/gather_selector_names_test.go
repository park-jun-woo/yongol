//ff:func feature=gen-splitter type=test control=sequence
//ff:what keepImport / selectorPkgIdent / gatherSelectorNames / isSourceFile
package splitter

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestGatherSelectorNames(t *testing.T) {
	src := "package p\nfunc f() { http.Get(\"u\"); fmt.Println(1) }"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	used := map[string]bool{}
	gatherSelectorNames(file.Decls[0], used)
	if !used["http"] || !used["fmt"] {
		t.Errorf("expected http and fmt collected, got %v", used)
	}
}
