//ff:func feature=cli-hash type=test control=sequence
//ff:what TestRun_DuplicateOp_ReturnsError — 중복 op 시 FT-01 에러 반환 확인

package clihash

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_DuplicateOp_ReturnsError(t *testing.T) {
	dir := t.TempDir()
	content := `tables:
  tasks: {}
features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
    table: tasks
  - op: CreateTask
    path: POST /tasks/v2
    desc: Create a task v2
    table: tasks
`
	if err := os.WriteFile(filepath.Join(dir, "features.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	err := Run(&buf, dir)
	if err == nil {
		t.Fatal("expected error for duplicate op, got nil")
	}
	if !strings.Contains(err.Error(), "[FT-01]") {
		t.Errorf("error should mention [FT-01]: %v", err)
	}

	// .yongol must not be created.
	if _, statErr := os.Stat(filepath.Join(dir, ".yongol")); statErr == nil {
		t.Error(".yongol should not be created when validation fails")
	}
}
