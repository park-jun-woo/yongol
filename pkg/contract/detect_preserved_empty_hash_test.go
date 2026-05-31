//ff:func feature=contract type=test control=sequence
//ff:what TestDetectPreservedBranches — read 에러/주석 없음/hash 빈값/일치(Untouched) 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

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
