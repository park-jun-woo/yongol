//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExtractExternalSymbolsQueries — server.Queries.XxxByY 호출이 SqlcQueries 에 수집됨

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbolsQueries(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nfunc Handler() {\n\tserver.Queries.GetUserByID(ctx, 1)\n\tserver.Queries.ListActiveSessions(ctx)\n}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sym, err := ExtractExternalSymbols(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]bool{"GetUserByID": true, "ListActiveSessions": true}
	if len(sym.SqlcQueries) != len(want) {
		t.Fatalf("queries: got %v want %v", sym.SqlcQueries, want)
	}
	for _, q := range sym.SqlcQueries {
		if !want[q] {
			t.Errorf("unexpected query: %q", q)
		}
	}
}
