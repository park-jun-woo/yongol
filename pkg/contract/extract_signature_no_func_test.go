//ff:func feature=contract type=test control=sequence
//ff:what TestExtractSignatureBranches — read/parse 에러·func 없음(zero)·init skip 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignature_NoFunc(t *testing.T) {
	path := filepath.Join(t.TempDir(), "typeonly.go")
	if err := os.WriteFile(path, []byte("package p\n\ntype T struct{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sig.Name != "" || len(sig.Params) != 0 {
		t.Errorf("expected zero FuncSignature, got %+v", sig)
	}
}
