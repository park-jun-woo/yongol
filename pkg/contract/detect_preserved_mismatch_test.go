//ff:func feature=contract type=test control=sequence
//ff:what test: TestDetectPreservedMismatch — saved hash 와 재계산 hash 가 다르면 StatePreserved

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectPreservedMismatch(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "edited.go")
	src := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\nfunc Demo() int { return 42 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	state, err := DetectPreserved(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != StatePreserved {
		t.Errorf("expected StatePreserved, got %v", state)
	}
}
