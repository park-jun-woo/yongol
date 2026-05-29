//ff:func feature=contract type=test control=sequence
//ff:what test: TestDetectPreservedNoAnnotation — //ff:checked 어노테이션이 없으면 StateNotApplicable

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectPreservedNoAnnotation(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "plain.go")
	src := "package demo\n\nfunc Demo() int { return 1 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	state, err := DetectPreserved(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != StateNotApplicable {
		t.Errorf("expected StateNotApplicable, got %v", state)
	}
}
