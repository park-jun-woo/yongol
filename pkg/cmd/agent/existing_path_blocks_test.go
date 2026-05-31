//ff:func feature=agent type=test control=sequence
//ff:what TestExistingPathBlocks — openapi.yaml의 path 블록 및 pathToOps 매핑 추출, 부재 시 nil 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExistingPathBlocks(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := `paths:
  /users:
    get:
      operationId: ListUsers
    post:
      operationId: CreateUser
`
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	paths, pathToOps := existingPathBlocks(dir)
	if _, ok := paths["/users"]; !ok {
		t.Errorf("paths missing /users: %v", paths)
	}
	ops := pathToOps["/users"]
	if len(ops) != 2 {
		t.Fatalf("pathToOps[/users] = %v, want 2 ops", ops)
	}
	got := map[string]bool{ops[0]: true, ops[1]: true}
	if !got["ListUsers"] || !got["CreateUser"] {
		t.Errorf("ops = %v, want ListUsers + CreateUser", ops)
	}

	// Missing file → nil maps.
	if p, m := existingPathBlocks(t.TempDir()); p != nil || m != nil {
		t.Errorf("missing file = %v, %v, want nil, nil", p, m)
	}
}
