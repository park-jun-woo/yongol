//ff:func feature=contract type=test control=sequence
//ff:what TestExtractSignatureBranches — read/parse 에러·func 없음(zero)·init skip 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSignature_ParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.go")
	if err := os.WriteFile(path, []byte("package p\nfunc ("), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ExtractSignature(path); err == nil {
		t.Fatal("want parse error")
	}
}
