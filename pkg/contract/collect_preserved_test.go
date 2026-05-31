//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectPreserved — 디렉토리 walk 시 preserved 파일만 반환
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectPreserved(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "svc")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}

	untouchedBody := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n")
	okHash := ComputeBodyHash(untouchedBody)
	okSrc := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=" + okHash + "\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n"
	preservedSrc := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\nfunc Demo() int { return 999 }\n"
	plainSrc := "package demo\n\nfunc Demo() int { return 1 }\n"

	if err := os.WriteFile(filepath.Join(sub, "ok.go"), []byte(okSrc), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "edited.go"), []byte(preservedSrc), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "plain.go"), []byte(plainSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := CollectPreserved(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 preserved file, got %d (%v)", len(got), got)
	}
	if filepath.Base(got[0]) != "edited.go" {
		t.Errorf("expected edited.go, got %s", got[0])
	}
}
