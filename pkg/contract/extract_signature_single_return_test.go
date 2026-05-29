//ff:func feature=contract type=test control=sequence
//ff:what test: TestExtractSignatureSingleReturn — 반환값 1개, error 없음 케이스

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignatureSingleReturn(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nfunc Count() int { return 0 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sig.Returns) != 1 || sig.Returns[0] != "int" {
		t.Errorf("returns: got %v want [int]", sig.Returns)
	}
	if sig.HasErr {
		t.Errorf("expected HasErr = false")
	}
}
