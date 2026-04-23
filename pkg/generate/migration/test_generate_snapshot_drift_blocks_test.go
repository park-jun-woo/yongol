//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestGenerate_SnapshotDrift_Blocks — 스냅샷 변조 시 MIG-006 로 차단
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGenerate_SnapshotDrift_Blocks(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "t.sql", `CREATE TABLE t (id BIGSERIAL PRIMARY KEY);`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	snapPath := filepath.Join(specsDir, "db", ".generated_schema.sql")
	data, _ := os.ReadFile(snapPath)
	tampered := append(data, []byte("\n-- tampered\n")...)
	if err := os.WriteFile(snapPath, tampered, 0644); err != nil {
		t.Fatal(err)
	}
	_, diags, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err == nil {
		t.Fatalf("expected Generate to fail after tamper, got nil")
	}
	foundMIG006 := false
	for _, d := range diags {
		if strings.Contains(d.Message, "MIG-006") {
			foundMIG006 = true
		}
	}
	if !foundMIG006 {
		t.Errorf("expected MIG-006 diag, got %+v", diags)
	}
}
