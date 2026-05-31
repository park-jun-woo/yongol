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

func TestRunRemove_ConfirmYesSkipsMissingSSaC(t *testing.T) {
	specs := setupSpecs(t)
	var out bytes.Buffer
	// Answer "y" -> proceeds; SSaC file doesn't exist -> skip branch.
	if err := RunRemove(&out, strings.NewReader("yes\n"), specs, []string{"DeleteTask"}, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skip") {
		t.Errorf("want skip message for missing ssac, got %q", out.String())
	}
	got, _ := os.ReadFile(filepath.Join(specs, "features.yaml"))
	if strings.Contains(string(got), "DeleteTask") {
		t.Error("DeleteTask should be removed from features.yaml")
	}
	if !strings.Contains(string(got), "CreateTask") {
		t.Error("CreateTask should remain")
	}
}
