//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV10PanicInitAllowed — init() 내부 panic 은 allowlist 대상

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV10PanicInitAllowed(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "bootstrap.go")
	writePreserved(t, p,
		"package service\n\nfunc init() { panic(\"cannot boot\") }\n")
	diags := prv10PreservedPanic([]string{p})
	if len(diags) != 0 {
		t.Fatalf("init() panic should be allowed, got %+v", diags)
	}
}
