//ff:func feature=migration type=test control=sequence
//ff:what TestGenerate_InitialMode — 스냅샷 없는 상태에서 0001_initial.sql + 스냅샷 헤더 생성
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGenerate_InitialMode(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL DEFAULT '');
`)
	res, diags, err := Generate(specsDir, artsDir, Options{
		YongolVersion: "v0.1.22",
		Now:           time.Date(2026, 4, 22, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("Generate: %v (diags=%+v)", err, diags)
	}
	if res.Mode != ModeInitial {
		t.Errorf("expected initial, got %v", res.Mode)
	}
	migPath := filepath.Join(artsDir, "db", "migrations", "0001_initial.sql")
	if _, err := os.Stat(migPath); err != nil {
		t.Errorf("0001_initial.sql missing: %v", err)
	}
	snapPath := filepath.Join(specsDir, "db", ".generated_schema.sql")
	if _, err := os.Stat(snapPath); err != nil {
		t.Errorf("snapshot missing: %v", err)
	}
	data, _ := os.ReadFile(snapPath)
	if !strings.HasPrefix(string(data), SnapshotHashHeaderPrefix) {
		t.Errorf("snapshot missing hash header:\n%s", data)
	}
}
