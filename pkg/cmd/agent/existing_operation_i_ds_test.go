//ff:func feature=agent type=test control=sequence
//ff:what TestExistingOperationIDs — openapi.yaml 의 operationId 수집, 부재 시 nil 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExistingOperationIDs(t *testing.T) {
	// Missing file returns nil.
	if got := existingOperationIDs(t.TempDir()); got != nil {
		t.Errorf("missing openapi.yaml = %v, want nil", got)
	}

	specs := t.TempDir()
	apiDir := filepath.Join(specs, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	doc := `paths:
  /users:
    get:
      operationId: ListUsers
    post:
      operationId: CreateUser
`
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(doc), 0o644); err != nil {
		t.Fatal(err)
	}
	ids := existingOperationIDs(specs)
	if !ids["ListUsers"] || !ids["CreateUser"] {
		t.Errorf("expected ListUsers and CreateUser, got %v", ids)
	}
	if len(ids) != 2 {
		t.Errorf("expected 2 ids, got %d: %v", len(ids), ids)
	}
}
