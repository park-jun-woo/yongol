//ff:func feature=gen-splitter type=test control=sequence
//ff:what keepImport / selectorPkgIdent / gatherSelectorNames / isSourceFile
package splitter

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsSourceFile(t *testing.T) {
	dir := t.TempDir()
	// plain source matching pattern with no ff header -> true
	src := filepath.Join(dir, "api.gen.go")
	if err := os.WriteFile(src, []byte("package p\nfunc f(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !isSourceFile(src, ToolOAPICodegen) {
		t.Errorf("plain .gen.go should be a source file")
	}
	// already-split output (ff:func header) -> false
	split := filepath.Join(dir, "f.gen.go")
	if err := os.WriteFile(split, []byte("//ff:func x\npackage p\nfunc f(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isSourceFile(split, ToolOAPICodegen) {
		t.Errorf("ff-annotated output should not be a source file")
	}
	// name not matching tool pattern -> false
	other := filepath.Join(dir, "plain.go")
	if err := os.WriteFile(other, []byte("package p\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isSourceFile(other, ToolOAPICodegen) {
		t.Errorf("non-matching name should not be a source file")
	}
}
