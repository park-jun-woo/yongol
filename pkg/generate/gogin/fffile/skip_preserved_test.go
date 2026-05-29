//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestWriteIfNotPreserved — 신규/untouched/preserved 3 케이스 write 여부 검증

package fffile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestWriteIfNotPreserved(t *testing.T) {
	dir := t.TempDir()

	// Case 1: file does not exist → must be written.
	newPath := filepath.Join(dir, "new.go")
	newBody := []byte("package x\n")
	if err := WriteIfNotPreserved(newPath, newBody); err != nil {
		t.Fatalf("new file write: %v", err)
	}
	if _, err := os.Stat(newPath); err != nil {
		t.Fatalf("expected new.go to exist: %v", err)
	}

	// Case 2: existing file that is untouched (hash matches body) →
	// overwrite permitted.
	body := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n")
	okHash := contract.ComputeBodyHash(body)
	okSrc := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=" + okHash + "\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n")
	untouched := filepath.Join(dir, "untouched.go")
	if err := os.WriteFile(untouched, okSrc, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := WriteIfNotPreserved(untouched, []byte("package demo\n// rewritten\n")); err != nil {
		t.Fatalf("untouched overwrite: %v", err)
	}
	out, _ := os.ReadFile(untouched)
	if string(out) != "package demo\n// rewritten\n" {
		t.Fatalf("expected overwrite on untouched, got %q", string(out))
	}

	// Case 3: existing preserved file (hash mismatch) → skipped.
	preservedSrc := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\nfunc Demo() int { return 999 }\n")
	preservedPath := filepath.Join(dir, "preserved.go")
	if err := os.WriteFile(preservedPath, preservedSrc, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := WriteIfNotPreserved(preservedPath, []byte("package demo\n// forbidden\n")); err != nil {
		t.Fatalf("preserved skip: %v", err)
	}
	out, _ = os.ReadFile(preservedPath)
	if string(out) != string(preservedSrc) {
		t.Fatalf("expected preserved body untouched, got %q", string(out))
	}
}
