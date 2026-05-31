//ff:func feature=contract type=test control=sequence
//ff:what TestExtractExternalSymbolsBranches — read/parse 에러·func 없는 파일(zero) 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbols_NoFunc(t *testing.T) {
	path := filepath.Join(t.TempDir(), "typeonly.go")
	if err := os.WriteFile(path, []byte("package svc\n\ntype T struct{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	sym, err := ExtractExternalSymbols(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sym.CallTargets) != 0 || len(sym.SqlcQueries) != 0 || len(sym.DDLFields) != 0 {
		t.Errorf("expected zero ExternalSymbols, got %+v", sym)
	}
}
