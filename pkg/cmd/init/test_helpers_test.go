//ff:func feature=cli-init type=test-helper control=sequence
//ff:what writeTempFeatures — writes a minimal features.yaml for test use

package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

const testFeaturesContent = `tables:
  workflows: {}
features:
  - op: CreateWorkflow
    path: POST /workflows
    desc: Create a new workflow
    table: workflows
  - op: GetWorkflow
    path: GET /workflows/{id}
    desc: Get workflow detail
    table: workflows
  - op: ListWorkflows
    path: GET /workflows
    desc: List all workflows
    table: workflows
`

// writeTempFeatures writes a minimal features.yaml in t.TempDir() and returns
// its absolute path.
func writeTempFeatures(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "features.yaml")
	if err := os.WriteFile(path, []byte(testFeaturesContent), 0o644); err != nil {
		t.Fatalf("write test features: %v", err)
	}
	return path
}
