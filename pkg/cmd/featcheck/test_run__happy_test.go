//ff:func feature=cli-featcheck type=test control=sequence
//ff:what TestRun_Happy

package featcheck

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun_Happy(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `tables:
  tasks: {}
features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
    table: tasks
  - op: GetTask
    path: GET /tasks/{id}
    desc: Get a task
    table: tasks
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	ff, diags, err := Run(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) != 0 {
		t.Errorf("want 0 diagnostics, got %d: %v", len(diags), diags)
	}
	if len(ff.Features) != 2 {
		t.Errorf("want 2 features, got %d", len(ff.Features))
	}
}
