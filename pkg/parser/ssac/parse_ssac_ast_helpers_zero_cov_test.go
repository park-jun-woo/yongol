//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseBatch_ZeroCov — ssac 파서 헬퍼를 이름으로 직접 호출해 커버 귀속
package ssac

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestParseSsacASTHelpers_ZeroCov(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "course.ssac", sampleSSaC, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse sample: %v", err)
	}
	imports := collectImports(f)
	_ = imports
	structs := collectStructs(f)
	if len(structs) == 0 {
		t.Error("expected at least one struct")
	}
	exerciseStructDecls(f)
	exerciseFuncDecls(fset, f, imports, structs)
	exerciseFileComments(fset, f)
}
