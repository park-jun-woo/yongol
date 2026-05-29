//ff:func feature=contract type=test control=sequence
//ff:what test: TestDetectPreservedUntouched — saved hash 와 재계산 hash 가 일치하면 StateUntouched

package contract

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDetectPreservedUntouched(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "untouched.go")
	bodyOnly := []byte("//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:what Demo — demo\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n")
	hash := ComputeBodyHash(bodyOnly)
	if hash == "" {
		t.Fatal("expected non-empty hash")
	}
	src := fmt.Sprintf(
		"//ff:func feature=demo type=handler control=sequence\n"+
			"//ff:what Demo — demo\n"+
			"//ff:checked llm=yongol-gen hash=%s\n"+
			"package demo\n\nfunc Demo() int { return 1 }\n",
		hash,
	)
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	state, err := DetectPreserved(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != StateUntouched {
		t.Errorf("expected StateUntouched, got %v", state)
	}
}
