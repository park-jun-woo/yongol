//ff:func feature=contract type=test control=sequence
//ff:what TestExtractExternalSymbolsBranches — read/parse 에러·func 없는 파일(zero) 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbols_ParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.go")
	if err := os.WriteFile(path, []byte("package svc\nfunc ("), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ExtractExternalSymbols(path); err == nil {
		t.Fatal("want parse error for malformed Go")
	}
}
