//ff:func feature=agent type=test control=sequence
//ff:what TestExistingPathBlocks — openapi.yaml의 path 블록 및 pathToOps 매핑 추출, 부재 시 nil 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExistingPathBlocksEdgeCases(t *testing.T) {
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
		if p, m := existingPathBlocks(write(t, "paths: [unclosed")); p != nil || m != nil {
			t.Errorf("malformed yaml = %v, %v, want nil, nil", p, m)
		}
	})

	t.Run("NoPaths", func(t *testing.T) {
		if p, m := existingPathBlocks(write(t, "openapi: 3.0.0\n")); p != nil || m != nil {
			t.Errorf("no paths = %v, %v, want nil, nil", p, m)
		}
	})

	t.Run("NonMapEntries", func(t *testing.T) {
		// Scalar path value → methods type-assert continue; scalar method detail
		// → detail type-assert continue. The /users path still yields its op.
		doc := `paths:
  /scalar: "string value"
  /users:
    get:
      operationId: ListUsers
    summary: "string detail"
`
		paths, pathToOps := existingPathBlocks(write(t, doc))
		if _, ok := paths["/scalar"]; !ok {
			t.Errorf("paths should still include /scalar key: %v", paths)
		}
		if ops := pathToOps["/users"]; len(ops) != 1 || ops[0] != "ListUsers" {
			t.Errorf("pathToOps[/users] = %v, want [ListUsers]", ops)
		}
		if _, ok := pathToOps["/scalar"]; ok {
			t.Errorf("scalar path should produce no ops, got %v", pathToOps["/scalar"])
		}
	})
}
