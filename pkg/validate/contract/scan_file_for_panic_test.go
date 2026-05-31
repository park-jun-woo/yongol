//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForPanic — 단일 파일 내 허용되지 않은 panic() 호출 수집 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForPanic(t *testing.T) {
	dir := t.TempDir()

	t.Run("panic in body flagged", func(t *testing.T) {
		p := filepath.Join(dir, "p.go")
		writePreserved(t, p, "package service\nfunc F() { panic(\"x\") }\n")
		if d := scanFileForPanic(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("panic in init allowed", func(t *testing.T) {
		p := filepath.Join(dir, "init.go")
		writePreserved(t, p, "package service\nfunc init() { panic(\"setup\") }\n")
		if d := scanFileForPanic(p); len(d) != 0 {
			t.Errorf("init panic should be allowed, got %+v", d)
		}
	})

	t.Run("parse error → nil", func(t *testing.T) {
		if d := scanFileForPanic(filepath.Join(dir, "missing.go")); d != nil {
			t.Errorf("expected nil for missing file, got %+v", d)
		}
	})
}
