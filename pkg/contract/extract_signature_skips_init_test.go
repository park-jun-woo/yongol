//ff:func feature=contract type=test control=sequence
//ff:what TestExtractSignatureBranches — read/parse 에러·func 없음(zero)·init skip 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignature_SkipsInit(t *testing.T) {
	path := filepath.Join(t.TempDir(), "withinit.go")
	src := "package p\n\nfunc init() {}\n\nfunc Real() {}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sig, err := ExtractSignature(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sig.Name != "Real" {
		t.Errorf("expected Real (init skipped), got %q", sig.Name)
	}
}
