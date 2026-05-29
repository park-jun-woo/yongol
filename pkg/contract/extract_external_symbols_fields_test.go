//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExtractExternalSymbolsFields — user.Email 같은 struct 필드 접근이 DDLFields 에 수집됨

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbolsFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nfunc Handler(user User) {\n\t_ = user.Email\n\t_ = user.name // lower — excluded\n\t_ = user.IsActive\n}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sym, err := ExtractExternalSymbols(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]bool{"user.Email": true, "user.IsActive": true}
	got := map[string]bool{}
	for _, f := range sym.DDLFields {
		got[f] = true
	}
	for k := range want {
		if !got[k] {
			t.Errorf("missing expected field %q in %v", k, sym.DDLFields)
		}
	}
	for k := range got {
		if !want[k] {
			t.Errorf("unexpected field %q", k)
		}
	}
}
