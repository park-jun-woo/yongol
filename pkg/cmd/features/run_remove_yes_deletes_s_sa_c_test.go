//ff:func feature=features type=test control=sequence
//ff:what TestRunRemove — 빈 ops/존재안함/abort/확인삭제/ssac-skip/--yes 성공 분기 검증
package features

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRemove_YesDeletesSSaC(t *testing.T) {
	specs := setupSpecs(t)
	// Create the SSaC file so the delete branch runs.
	ssacDir := filepath.Join(specs, "service", "task")
	if err := os.MkdirAll(ssacDir, 0o755); err != nil {
		t.Fatal(err)
	}
	ssacFile := filepath.Join(ssacDir, "DeleteTask.ssac")
	if err := os.WriteFile(ssacFile, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), specs, []string{"DeleteTask"}, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := os.Stat(ssacFile); !os.IsNotExist(err) {
		t.Error("expected SSaC file to be deleted")
	}
	if !strings.Contains(out.String(), "1 feature(s) removed, 1 SSaC file(s) deleted") {
		t.Errorf("want summary, got %q", out.String())
	}
	// .yongol hash written.
	if _, err := os.Stat(filepath.Join(specs, ".yongol")); err != nil {
		t.Errorf("expected .yongol hash: %v", err)
	}
}
