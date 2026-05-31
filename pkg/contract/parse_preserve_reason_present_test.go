//ff:func feature=contract type=test control=sequence
//ff:what TestParsePreserveReasonBranches — read 에러 / reason 추출 성공 분기 검증
package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePreserveReason_Present(t *testing.T) {
	path := filepath.Join(t.TempDir(), "p.go")
	src := "//ff:preserve reason=\"manual tuning of retry loop\"\n" +
		"package demo\n\nfunc Demo() int { return 1 }\n"
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := ParsePreserveReason(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "manual tuning of retry loop" {
		t.Errorf("reason = %q, want %q", got, "manual tuning of retry loop")
	}
}
