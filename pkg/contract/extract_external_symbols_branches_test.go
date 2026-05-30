//ff:func feature=contract type=test control=iteration dimension=1
//ff:what TestExtractExternalSymbolsBranches — read/parse 에러·func 없는 파일(zero) 분기 검증

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbols_ReadError(t *testing.T) {
	if _, err := ExtractExternalSymbols(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error")
	}
}

func TestExtractExternalSymbols_ParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.go")
	if err := os.WriteFile(path, []byte("package svc\nfunc ("), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ExtractExternalSymbols(path); err == nil {
		t.Fatal("want parse error for malformed Go")
	}
}

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

func TestExtractExternalSymbols_SkipsInit(t *testing.T) {
	// init func is skipped; the following func is the one walked.
	path := filepath.Join(t.TempDir(), "withinit.go")
	src := "package svc\n\nfunc init() { setup.Run() }\n\nfunc H() { billing.Charge(ctx) }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sym, err := ExtractExternalSymbols(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	found := false
	for _, c := range sym.CallTargets {
		if c == "billing.Charge" {
			found = true
		}
		if c == "setup.Run" {
			t.Errorf("init func should be skipped, got %q", c)
		}
	}
	if !found {
		t.Errorf("expected billing.Charge from H(), got %v", sym.CallTargets)
	}
}
