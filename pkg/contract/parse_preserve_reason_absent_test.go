//ff:func feature=contract type=test control=sequence
//ff:what test: TestParsePreserveReasonAbsent — //ff:preserve 어노테이션이 없으면 빈 문자열

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePreserveReasonAbsent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "p.go")
	src := "package demo\n\nfunc Demo() int { return 1 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := ParsePreserveReason(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Errorf("expected empty reason, got %q", got)
	}
}
