//ff:func feature=cli type=test control=sequence
//ff:what test: TestCountPreserved — validates split count of files with and without a reason comment

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountPreserved(t *testing.T) {
	dir := t.TempDir()
	withReasonPath := filepath.Join(dir, "a.go")
	withoutReasonPath := filepath.Join(dir, "b.go")
	missingPath := filepath.Join(dir, "missing.go") // not created on purpose

	withReasonSrc := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:preserve reason=\"custom batch\"\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\nfunc Demo() int { return 999 }\n"
	withoutReasonSrc := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\nfunc Demo() int { return 999 }\n"
	if err := os.WriteFile(withReasonPath, []byte(withReasonSrc), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(withoutReasonPath, []byte(withoutReasonSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	wr, wo := countPreserved([]string{withReasonPath, withoutReasonPath, missingPath})
	if wr != 1 {
		t.Errorf("expected 1 withReason, got %d", wr)
	}
	if wo != 2 {
		t.Errorf("expected 2 withoutReason (incl. missing), got %d", wo)
	}
}
