//ff:func feature=cli-featcheck type=test control=iteration dimension=1
//ff:what TestRun_DuplicateOp_FT01

package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_DuplicateOp_FT01(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
  - op: CreateTask
    path: POST /tasks/v2
    desc: Create a task v2
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	_, diags, err := Run(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) == 0 {
		t.Fatal("want diagnostics for duplicate op, got 0")
	}
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "[FT-01]") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("want [FT-01] diagnostic, got %v", diags)
	}
}
