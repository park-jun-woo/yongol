//ff:what Load — features.yaml 정상 파싱 테스트
package features

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_Happy(t *testing.T) {
	dir := t.TempDir()
	data := `features:
  - op: CreateWorkflow
    path: POST /workflows
    desc: Create a new workflow
  - op: GetWorkflow
    path: GET /workflows/{id}
    desc: Get a workflow
`
	if err := os.WriteFile(filepath.Join(dir, "features.yaml"), []byte(data), 0644); err != nil {
		t.Fatal(err)
	}
	feats, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(feats) != 2 {
		t.Fatalf("want 2 features, got %d", len(feats))
	}
	if feats[0].Op != "CreateWorkflow" {
		t.Errorf("want CreateWorkflow, got %s", feats[0].Op)
	}
	if feats[0].Line == 0 {
		t.Error("expected non-zero line for first feature")
	}
	if feats[1].Path != "GET /workflows/{id}" {
		t.Errorf("want GET /workflows/{id}, got %s", feats[1].Path)
	}
}
