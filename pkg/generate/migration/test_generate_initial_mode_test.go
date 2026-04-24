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

-- @sentinel
INSERT INTO users (id, name) VALUES (0, 'nobody') ON CONFLICT DO NOTHING;
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
	upData, err := os.ReadFile(upPath)
	if err != nil {
		t.Fatalf("0001_initial.up.sql missing: %v", err)
	}
	up := string(upData)
	// Sentinel INSERT must be present and appear AFTER CREATE TABLE.
	if !strings.Contains(up, "INSERT INTO users") {
		t.Errorf("initial migration missing sentinel INSERT:\n%s", up)
	}
	createIdx := strings.Index(up, "CREATE TABLE")
	insertIdx := strings.Index(up, "INSERT INTO users")
	if createIdx < 0 || insertIdx < 0 || createIdx >= insertIdx {
		t.Errorf("sentinel INSERT must follow CREATE TABLE; createIdx=%d insertIdx=%d\n%s",
			createIdx, insertIdx, up)
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
	if !strings.Contains(string(data), "INSERT INTO users") {
		t.Errorf("snapshot missing sentinel INSERT body:\n%s", data)
	}
}
