//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestCopyPrune_ZeroCov — copyHurlFile / pruneOrphans 정상·에러·no-op 분기 직접 호출
package hurl_mirror

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyHurlFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.hurl")
	if err := os.WriteFile(src, []byte("GET http://x\nHTTP 200\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	dst := filepath.Join(dir, "nested", "out.hurl")
	if err := copyHurlFile(src, dst); err != nil {
		t.Fatalf("copyHurlFile: %v", err)
	}
	got, err := os.ReadFile(dst)
	if err != nil || len(got) == 0 {
		t.Fatalf("dst not written: %v", err)
	}
	// missing src → error.
	if err := copyHurlFile(filepath.Join(dir, "nope.hurl"), filepath.Join(dir, "x.hurl")); err == nil {
		t.Errorf("expected error for missing src")
	}
}
