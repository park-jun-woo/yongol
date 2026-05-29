//ff:func feature=features type=test control=sequence
//ff:what Load — path 필드 누락 시 에러 진단 테스트
package features

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_MissingPath(t *testing.T) {
	dir := t.TempDir()
	data := `features:
  - op: CreateWorkflow
    desc: Create a new workflow
`
	if err := os.WriteFile(filepath.Join(dir, "features.yaml"), []byte(data), 0644); err != nil {
		t.Fatal(err)
	}
	_, diags := Load(dir)
	if len(diags) == 0 {
		t.Fatal("expected diag for missing path")
	}
}
