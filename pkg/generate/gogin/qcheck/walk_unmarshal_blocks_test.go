//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWalkUnmarshalBlocks — 전 함수 순회로 미가드 Unmarshal DF-01 수집 + 비함수/본문없음 스킵
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkUnmarshalBlocks(t *testing.T) {
	src := `package x

func decl()

func ok(b []byte, v any) {
	if err := json.Unmarshal(b, v); err != nil { return }
}

func bad(b []byte, v any) {
	_ = yaml.Unmarshal(b, v)
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	findings := walkUnmarshalBlocks(file, fset)
	if len(findings) != 1 {
		t.Fatalf("want 1 DF-01 finding (yaml in bad()), got %d: %+v", len(findings), findings)
	}
	if findings[0].Detail != "yaml.Unmarshal" {
		t.Errorf("unexpected finding: %+v", findings[0])
	}
}
