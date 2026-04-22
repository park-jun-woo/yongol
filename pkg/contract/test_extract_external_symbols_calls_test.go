//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExtractExternalSymbolsCalls — pkg.Func(...) 호출이 CallTargets 에 수집됨

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractExternalSymbolsCalls(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.go")
	src := "package svc\n\nfunc Handler() {\n\tbilling.HoldEscrow(ctx, 100)\n\tnotify.Send(ctx, \"ok\")\n}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	sym, err := ExtractExternalSymbols(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]bool{"billing.HoldEscrow": true, "notify.Send": true}
	if len(sym.CallTargets) != len(want) {
		t.Fatalf("calls: got %v want %v", sym.CallTargets, want)
	}
	for _, c := range sym.CallTargets {
		if !want[c] {
			t.Errorf("unexpected call: %q", c)
		}
	}
}
