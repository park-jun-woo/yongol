//ff:func feature=contract type=test control=iteration dimension=1
//ff:what TestExtractSignatureBranches — read/parse 에러·func 없음(zero)·init skip 분기 검증

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignature_ReadError(t *testing.T) {
	if _, err := ExtractSignature(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error")
	}
}

func TestExtractSignature_ParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.go")
	if err := os.WriteFile(path, []byte("package p\nfunc ("), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ExtractSignature(path); err == nil {
		t.Fatal("want parse error")
	}
}

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
