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
