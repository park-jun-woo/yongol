//ff:func feature=agent type=test control=sequence
//ff:what TestScaffold — 테이블 없음 skip / 테이블 존재+미지원 backend → DDL phase 에러 분기 검증

package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldNoTables(t *testing.T) {
	var out bytes.Buffer
	// nil FeaturesFile → skip.
	if err := scaffold(t.TempDir(), nil, nil, Config{}, &out); err != nil {
		t.Fatalf("nil ff: unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "scaffold skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}

	// FeaturesFile with no tables → skip.
	out.Reset()
	if err := scaffold(t.TempDir(), &features.FeaturesFile{}, nil, Config{}, &out); err != nil {
		t.Fatalf("empty ff: unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "scaffold skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}

func TestScaffoldDDLError(t *testing.T) {
	// One table + an unsupported backend makes the DDL phase's LLM call fail, so
	// scaffold returns a wrapped "scaffold DDL" error.
	ff := &features.FeaturesFile{
		Tables: map[string]features.TableDef{
			"users": {},
		},
	}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	err := scaffold(t.TempDir(), ff, nil, cfg, &out)
	if err == nil {
		t.Fatal("expected DDL phase error")
	}
	if !strings.Contains(err.Error(), "scaffold DDL") {
		t.Errorf("expected scaffold DDL error, got: %v", err)
	}
}

func TestScaffoldSQLcError(t *testing.T) {
	// Pre-create the DDL file so the DDL phase skips, then the sqlc phase's LLM
	// call fails for the unsupported backend, yielding a "scaffold sqlc" error.
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
	err := scaffold(dir, ff, nil, cfg, &out)
	if err == nil || !strings.Contains(err.Error(), "scaffold sqlc") {
		t.Fatalf("expected scaffold sqlc error, got: %v", err)
	}
}

func TestScaffoldFullPipelineNoLLM(t *testing.T) {
	// One table (no states) and ZERO features: pre-creating the DDL and sqlc
	// files makes both LLM phases skip, while the OpenAPI/SSaC/Rego/stateDiagram
	// phases short-circuit on the empty feature list. The full pipeline runs to
	// the completion summary with no network calls.
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	queriesDir := filepath.Join(dbDir, "queries")
	if err := os.MkdirAll(queriesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte("CREATE TABLE users();\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(queriesDir, "users.sql"), []byte("-- name: GetUser :one\nSELECT 1;\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Pre-create api/openapi.yaml so the post-OpenAPI ReadFile branch (loading
	// the doc for SSaC prompts) is exercised.
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte("paths: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffold(dir, ff, nil, cfg, &out); err != nil {
		t.Fatalf("expected full pipeline to complete, got: %v", err)
	}
	if !strings.Contains(out.String(), "scaffold complete") {
		t.Errorf("expected completion summary, got: %q", out.String())
	}
}
