//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what zz_zerocov_test — gogin.WriteManyFiles 0% 커버리지 단위 테스트
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteManyFiles_ZeroCov(t *testing.T) {
	// Empty input → empty set, no dir created.
	dir := filepath.Join(t.TempDir(), "sub")
	written, err := WriteManyFiles(dir, map[string]string{})
	if err != nil || len(written) != 0 {
		t.Fatalf("empty: written=%v err=%v", written, err)
	}
	if _, statErr := os.Stat(dir); !os.IsNotExist(statErr) {
		t.Error("dir should not be created for empty input")
	}

	// Non-empty → writes files, .go gets checked-line injected.
	dir2 := filepath.Join(t.TempDir(), "out")
	files := map[string]string{
		"handler.go": "package main\n\nfunc Foo() {}\n",
		"notes.txt":  "hello",
	}
	written, err = WriteManyFiles(dir2, files)
	if err != nil {
		t.Fatalf("write: %v", err)
	}
	if !written["handler.go"] || !written["notes.txt"] {
		t.Errorf("written set incomplete: %v", written)
	}
	for name := range files {
		if _, err := os.Stat(filepath.Join(dir2, name)); err != nil {
			t.Errorf("%s not on disk: %v", name, err)
		}
	}
}
