//ff:func feature=generate type=test control=sequence
//ff:what Generate 오케스트레이터의 migration no-op + backend 에러 경로 검증

package generate

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_MigrationErrorPropagates(t *testing.T) {
	// A SpecsDir with a malformed DDL file makes the migration step fail,
	// which Generate must wrap and return before backend codegen.
	specsDir := t.TempDir()
	ddlDir := filepath.Join(specsDir, "db")
	if err := os.MkdirAll(ddlDir, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(ddlDir, "schema.sql"), []byte("CREATE TABLE ("), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	fs := &yongol.Fullstack{SpecsDir: specsDir}
	err := Generate(fs, t.TempDir(), GoGin, React)
	if err == nil || !strings.Contains(err.Error(), "migration") {
		t.Fatalf("expected migration error, got: %v", err)
	}
}

func TestGenerate_UnknownBackendErrors(t *testing.T) {
	// Empty SpecsDir makes runMigrationStep a no-op; an unknown backend then
	// surfaces a wrapped "backend:" error before frontend/hurl/opa steps run.
	fs := &yongol.Fullstack{SpecsDir: ""}
	err := Generate(fs, t.TempDir(), BackendType("does-not-exist"), FrontendType("react"))
	if err == nil || !strings.Contains(err.Error(), "backend:") {
		t.Fatalf("expected backend error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "unknown backend") {
		t.Errorf("expected 'unknown backend' in error, got: %v", err)
	}
}
