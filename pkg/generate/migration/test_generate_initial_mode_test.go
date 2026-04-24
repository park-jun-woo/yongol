//ff:func feature=migration type=test control=sequence
//ff:what TestGenerate_InitialMode — 스냅샷 없는 상태에서 0001_initial.up.sql + down stub + 스냅샷 헤더 생성
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
	if res.MigrationFile != "0001_initial.up.sql" {
		t.Errorf("expected MigrationFile=0001_initial.up.sql, got %q", res.MigrationFile)
	}
	upPath := filepath.Join(artsDir, "db", "migrations", "0001_initial.up.sql")
	if _, err := os.Stat(upPath); err != nil {
		t.Errorf("0001_initial.up.sql missing: %v", err)
	}
	downPath := filepath.Join(artsDir, "db", "migrations", "0001_initial.down.sql")
	if _, err := os.Stat(downPath); err != nil {
		t.Errorf("0001_initial.down.sql missing: %v", err)
	}
	downData, _ := os.ReadFile(downPath)
	if !strings.Contains(string(downData), "Down stub — intentionally empty.") {
		t.Errorf("down stub missing marker line:\n%s", downData)
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
