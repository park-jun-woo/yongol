//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Generate 엔트리포인트 e2e — initial / incremental / noop / MIG-006 drift
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeSpec(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGenerate_InitialMode(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL DEFAULT '');
`)
	res, diags, err := Generate(specsDir, artsDir, Options{
		YongolVersion: "v0.1.22",
		Now:            time.Date(2026, 4, 22, 0, 0, 0, 0, time.UTC),
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
	// Snapshot should have YONGOL_SCHEMA_HASH header
	data, _ := os.ReadFile(snapPath)
	if !strings.HasPrefix(string(data), SnapshotHashHeaderPrefix) {
		t.Errorf("snapshot missing hash header:\n%s", data)
	}
}

func TestGenerate_IncrementalMode(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL DEFAULT '');
`)
	// First call: creates initial
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	// Modify DDL — add column with DEFAULT so no MIG-002.
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL DEFAULT '', age INTEGER NOT NULL DEFAULT 0);
`)
	res, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err != nil {
		t.Fatalf("incremental: %v", err)
	}
	if res.Mode != ModeIncremental {
		t.Errorf("expected incremental, got %v", res.Mode)
	}
	if !strings.Contains(res.MigrationFile, "0002_") {
		t.Errorf("expected 0002_* file, got %q", res.MigrationFile)
	}
}

func TestGenerate_NoopMode(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "users.sql", `
CREATE TABLE users (id BIGSERIAL PRIMARY KEY);
`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	// Second run with no DDL change.
	res, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err != nil {
		t.Fatalf("noop: %v", err)
	}
	if res.Mode != ModeNoop {
		t.Errorf("expected noop, got %v", res.Mode)
	}
	// No 0002_* file should exist.
	entries, _ := os.ReadDir(filepath.Join(artsDir, "db", "migrations"))
	if len(entries) != 1 {
		t.Errorf("expected exactly 1 migration file after noop, got %d: %+v", len(entries), entries)
	}
}

func TestGenerate_SnapshotDrift_Blocks(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "t.sql", `CREATE TABLE t (id BIGSERIAL PRIMARY KEY);`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	// Tamper the snapshot.
	snapPath := filepath.Join(specsDir, "db", ".generated_schema.sql")
	data, _ := os.ReadFile(snapPath)
	// Append garbage after the hash header to invalidate the hash.
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

func TestGenerate_NotNullWithoutBackfill_Blocks(t *testing.T) {
	specsDir := t.TempDir()
	artsDir := t.TempDir()
	writeSpec(t, filepath.Join(specsDir, "db"), "t.sql", `CREATE TABLE t (id BIGSERIAL PRIMARY KEY);`)
	if _, _, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()}); err != nil {
		t.Fatalf("initial: %v", err)
	}
	// Add NOT NULL without default or backfill.
	writeSpec(t, filepath.Join(specsDir, "db"), "t.sql", `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, active BOOLEAN NOT NULL);`)
	_, diags, err := Generate(specsDir, artsDir, Options{YongolVersion: "v1", Now: time.Now().UTC()})
	if err == nil {
		t.Fatalf("expected block, got nil")
	}
	foundMIG002 := false
	for _, d := range diags {
		if strings.Contains(d.Message, "MIG-002") {
			foundMIG002 = true
		}
	}
	if !foundMIG002 {
		t.Errorf("expected MIG-002, got %+v", diags)
	}
}
