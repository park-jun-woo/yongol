//ff:func feature=contract type=test control=sequence
//ff:what TestDetectPreservedBranches — read 에러/주석 없음/hash 빈값/일치(Untouched) 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

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
