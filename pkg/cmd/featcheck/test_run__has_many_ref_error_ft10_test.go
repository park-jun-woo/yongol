//ff:func feature=cli-featcheck type=test control=iteration dimension=1
//ff:what TestRun_HasManyRefError_FT10

package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_HasManyRefError_FT10(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	content := `tables:
  projects:
    has_many:
      - tasks
      - ghost_table
features:
  - op: CreateProject
    path: POST /projects
    desc: Create a project
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	_, diags, err := Run(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) == 0 {
		t.Fatal("want diagnostics for has_many ref error, got 0")
	}
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "[FT-10]") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("want [FT-10] diagnostic, got %v", diags)
	}
}
