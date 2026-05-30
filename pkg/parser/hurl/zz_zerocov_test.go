//ff:func feature=orchestrator type=test
//ff:what zz_zerocov_test — hurl.CollectFiles 0% 커버리지 단위 테스트
package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	write := func(name string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("a.hurl")
	write("b.hurl")
	write("c.txt") // ignored
	if err := os.Mkdir(filepath.Join(dir, "nested.hurl"), 0o755); err != nil {
		t.Fatal(err) // dir with .hurl suffix is skipped
	}

	files := CollectFiles(dir)
	if len(files) != 2 {
		t.Fatalf("expected 2 .hurl files, got %d: %v", len(files), files)
	}
	for _, f := range files {
		if filepath.Ext(f) != ".hurl" {
			t.Errorf("unexpected file: %s", f)
		}
	}

	// Missing dir → nil.
	if got := CollectFiles(filepath.Join(dir, "absent")); got != nil {
		t.Errorf("missing dir should return nil, got %v", got)
	}
}
