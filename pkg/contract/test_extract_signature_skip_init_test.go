//ff:func feature=contract type=test control=sequence
//ff:what test: TestExtractSignatureSkipInit — init 함수는 건너뛰고 첫 일반 함수의 시그니처를 반환

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignatureSkipInit(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nfunc init() {}\n\nfunc Real() {}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sig.Name != "Real" {
		t.Errorf("got %q want Real", sig.Name)
	}
}
