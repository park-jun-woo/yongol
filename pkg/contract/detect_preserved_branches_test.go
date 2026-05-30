//ff:func feature=contract type=test control=sequence
//ff:what TestDetectPreservedBranches — read 에러/주석 없음/hash 빈값/일치(Untouched) 분기 검증

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectPreserved_ReadError(t *testing.T) {
	if _, err := DetectPreserved(filepath.Join(t.TempDir(), "nope.go")); err == nil {
		t.Fatal("want read error for missing file")
	}
}

func TestDetectPreserved_NoAnnotation(t *testing.T) {
	path := filepath.Join(t.TempDir(), "plain.go")
	src := "package demo\n\nfunc Demo() int { return 1 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	state, err := DetectPreserved(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != StateNotApplicable {
		t.Errorf("want StateNotApplicable, got %v", state)
	}
}

func TestDetectPreserved_EmptyHash(t *testing.T) {
	// Has a checked annotation but no func decl -> ComputeBodyHash == "" -> NotApplicable.
	path := filepath.Join(t.TempDir(), "nofunc.go")
	src := "//ff:checked llm=yongol-gen hash=deadbeef\npackage demo\n\ntype T struct{}\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	state, err := DetectPreserved(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != StateNotApplicable {
		t.Errorf("want StateNotApplicable (empty hash), got %v", state)
	}
}

func TestDetectPreserved_Untouched(t *testing.T) {
	dir := t.TempDir()
	body := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n")
	h := ComputeBodyHash(body)
	src := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=" + h + "\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n"
	path := filepath.Join(dir, "ok.go")
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	state, err := DetectPreserved(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != StateUntouched {
		t.Errorf("want StateUntouched, got %v", state)
	}
}
