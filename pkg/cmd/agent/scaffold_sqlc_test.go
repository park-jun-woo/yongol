//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestScaffoldSQLc — 테이블 없음 nil / 기존파일 skip / 미지원 backend LLM 에러 / mkdir 에러 분기 검증

package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSQLcNoTables(t *testing.T) {
	var out bytes.Buffer
	if err := scaffoldSQLc(t.TempDir(), &features.FeaturesFile{}, nil, Config{}, &out); err != nil {
		t.Fatalf("no tables: unexpected error: %v", err)
	}
}

func TestScaffoldSQLcSkipExisting(t *testing.T) {
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	queriesDir := filepath.Join(dbDir, "queries")
	if err := os.MkdirAll(queriesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte("CREATE TABLE users();\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(queriesDir, "users.sql"), []byte("-- name: G :one\nSELECT 1;\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	if err := scaffoldSQLc(dir, ff, nil, Config{}, &out); err != nil {
		t.Fatalf("skip-existing: unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}

func TestScaffoldSQLcLLMError(t *testing.T) {
	// DDL present but queries missing → scaffoldSQLcTable's LLM call fails.
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte("CREATE TABLE users();\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffoldSQLc(dir, ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldSQLcTable")
	}
}

func TestScaffoldSQLcMkdirError(t *testing.T) {
	// Creating db/queries as a regular file makes os.MkdirAll fail.
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "queries"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	if err := scaffoldSQLc(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when db/queries is a file")
	}
}
