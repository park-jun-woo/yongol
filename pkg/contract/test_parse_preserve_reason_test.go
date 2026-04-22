//ff:func feature=contract type=test control=sequence
//ff:what test: TestParsePreserveReasonPresent — reason 어노테이션이 있을 때 문자열 추출

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePreserveReasonPresent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "p.go")
	src := "//ff:func feature=demo type=handler control=sequence\n" +
		"//ff:preserve reason=\"custom logic for premium users\"\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := ParsePreserveReason(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "custom logic for premium users" {
		t.Errorf("unexpected reason: %q", got)
	}
}
