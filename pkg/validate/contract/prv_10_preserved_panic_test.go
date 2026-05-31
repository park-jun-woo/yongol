//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV10PreservedPanic — preserved 파일 목록 panic 오케스트레이션 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV10PreservedPanic(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "ok.go")
	writePreserved(t, good, "package service\nfunc F() error { return nil }\n")
	bad := filepath.Join(dir, "bad.go")
	writePreserved(t, bad, "package service\nfunc G() { panic(\"x\") }\n")

	diags := prv10PreservedPanic([]string{good, bad})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d (%+v)", len(diags), diags)
	}
	if diags[0].File != bad {
		t.Errorf("expected diag on bad file, got %q", diags[0].File)
	}
}
