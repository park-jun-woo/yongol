//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestCollectPreserved — 디렉토리 walk 시 preserved 파일만 반환
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectPreserved_SkipsDirsAndNonGo(t *testing.T) {
	root := t.TempDir()
	preservedSrc := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\nfunc Demo() int { return 999 }\n")

	// Preserved .go inside a dotted dir, a vendor dir, and node_modules — all skipped.
	for _, d := range []string{".git", "vendor", "node_modules"} {
		dir := filepath.Join(root, d)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "x.go"), preservedSrc, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	// A non-Go file is ignored.
	if err := os.WriteFile(filepath.Join(root, "readme.txt"), []byte("hi"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := CollectPreserved(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected 0 preserved (all skipped), got %v", got)
	}
}
