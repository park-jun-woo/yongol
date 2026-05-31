//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what TestCopyPrune_ZeroCov — copyHurlFile / pruneOrphans 정상·에러·no-op 분기 직접 호출
package hurl_mirror

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPruneOrphans_ZeroCov(t *testing.T) {
	// missing dstRoot → no-op nil.
	if err := pruneOrphans(filepath.Join(t.TempDir(), "absent"), map[string]struct{}{}); err != nil {
		t.Errorf("missing dstRoot should be nil, got %v", err)
	}

	root := t.TempDir()
	keep := filepath.Join(root, "keep.hurl")
	orphan := filepath.Join(root, "sub", "orphan.hurl")
	other := filepath.Join(root, "readme.txt")
	for _, p := range []string{keep, orphan, other} {
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	kept := map[string]struct{}{"keep.hurl": {}}
	if err := pruneOrphans(root, kept); err != nil {
		t.Fatalf("pruneOrphans: %v", err)
	}
	if _, err := os.Stat(keep); err != nil {
		t.Errorf("kept file was removed")
	}
	if _, err := os.Stat(orphan); !os.IsNotExist(err) {
		t.Errorf("orphan not pruned")
	}
	if _, err := os.Stat(other); err != nil {
		t.Errorf("non-hurl file removed")
	}

	// dstRoot is a regular file → no-op.
	f := filepath.Join(root, "keep.hurl")
	if err := pruneOrphans(f, kept); err != nil {
		t.Errorf("file dstRoot should be no-op, got %v", err)
	}
}
