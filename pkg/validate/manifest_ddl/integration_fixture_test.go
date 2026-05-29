//ff:func feature=validate type=test-helper control=sequence topic=manifest-infra
//ff:what writeXDNFixture — temp specs 디렉토리에 manifest.yaml + db/users.sql 작성

package manifest_ddl

import (
	"os"
	"path/filepath"
	"testing"
)

// writeXDNFixture writes a manifest.yaml + db/users.sql pair under a fresh
// t.TempDir, returning the specs root. Callers tweak the contents per
// test by passing manifestBody and ddlBody.
func writeXDNFixture(t *testing.T, manifestBody, ddlBody string) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "manifest.yaml"), []byte(manifestBody), 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	dbDir := filepath.Join(root, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir db: %v", err)
	}
	if ddlBody != "" {
		if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte(ddlBody), 0o644); err != nil {
			t.Fatalf("write users.sql: %v", err)
		}
	}
	return root
}
