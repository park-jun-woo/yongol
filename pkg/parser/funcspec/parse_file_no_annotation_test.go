//ff:func feature=funcspec type=parser control=sequence
//ff:what ParseFile 어노테이션 없는 파일 테스트 — @func 없으면 nil 반환

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileNoAnnotation(t *testing.T) {
	dir := t.TempDir()
	src := `package foo

func Foo() {}
`
	path := filepath.Join(dir, "foo.go")
	os.WriteFile(path, []byte(src), 0644)

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if spec != nil {
		t.Error("expected nil for file without @func annotation")
	}
}
