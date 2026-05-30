//ff:func feature=agent type=test control=iteration dimension=3
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

func TestExistingOperationIDsEdgeCases(t *testing.T) {
	write := func(t *testing.T, content string) string {
		t.Helper()
		specs := t.TempDir()
		apiDir := filepath.Join(specs, "api")
		if err := os.MkdirAll(apiDir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return specs
	}

	t.Run("MalformedYAML", func(t *testing.T) {
		// Broken YAML → Unmarshal error → nil.
		if got := existingOperationIDs(write(t, "paths: [unclosed")); got != nil {
			t.Errorf("malformed yaml = %v, want nil", got)
		}
	})

	t.Run("NoPaths", func(t *testing.T) {
		// Valid YAML without a paths map → doc.Paths nil → nil.
		if got := existingOperationIDs(write(t, "openapi: 3.0.0\n")); got != nil {
			t.Errorf("no paths = %v, want nil", got)
		}
	})

	t.Run("MethodsNotMap", func(t *testing.T) {
		// A path whose value is a scalar (not a map) hits the methods type-assert
		// continue; a $ref string under a method hits the detail type-assert
		// continue. The well-formed /users path still yields its operationId.
		doc := `paths:
  /scalar: "just a string"
  /users:
    get:
      operationId: ListUsers
    summary: "a string detail, not a map"
`
		ids := existingOperationIDs(write(t, doc))
		if !ids["ListUsers"] {
			t.Errorf("expected ListUsers despite non-map siblings, got %v", ids)
		}
		if len(ids) != 1 {
			t.Errorf("expected exactly 1 id, got %d: %v", len(ids), ids)
		}
	})
}
