//ff:func feature=contract type=test control=sequence
//ff:what test: TestExtractSignatureNoParams — 파라미터·반환값 없는 함수 시그니처

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignatureNoParams(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nfunc Run() {}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sig.Name != "Run" {
		t.Errorf("name: got %q want Run", sig.Name)
	}
	if len(sig.Params) != 0 {
		t.Errorf("expected no params, got %v", sig.Params)
	}
	if len(sig.Returns) != 0 {
		t.Errorf("expected no returns, got %v", sig.Returns)
	}
	if sig.HasErr {
		t.Errorf("expected HasErr = false")
	}
}
