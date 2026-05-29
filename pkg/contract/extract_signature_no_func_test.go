//ff:func feature=contract type=test control=sequence
//ff:what test: TestExtractSignatureNoFunc — type-only 파일은 zero FuncSignature 를 반환

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignatureNoFunc(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "types.go")
	src := "package types\n\ntype User struct { ID int64 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sig.Name != "" || len(sig.Params) != 0 || len(sig.Returns) != 0 {
		t.Errorf("expected zero FuncSignature, got %+v", sig)
	}
}
