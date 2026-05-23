//ff:func feature=cli-featcheck type=test control=sequence
//ff:what TestRun_MissingOp_ReturnsError

package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_MissingOp_ReturnsError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `features:
  - op: ""
    path: POST /tasks
    desc: Create a task
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	_, _, err := Run(path)
	if err == nil {
		t.Fatal("want error for missing op, got nil")
	}
	if !strings.Contains(err.Error(), "missing required field 'op'") {
		t.Errorf("unexpected error: %v", err)
	}
}
