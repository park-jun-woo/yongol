//ff:func feature=gen-gogin type=test control=branch topic=request-validator
//ff:what TestGenerate — nil fs early-return + copy 에러 + 전체 성공 경로 검증

package middleware

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_NilFullstack(t *testing.T) {
	if err := Generate(nil, prepared.State{}, t.TempDir()); err != nil {
		t.Errorf("nil fs should return nil, got: %v", err)
	}
}

func TestGenerate_CopyOpenAPIError(t *testing.T) {
	// SpecsDir/api/openapi.yaml does not exist -> copyFile open error.
	fs := &yongol.Fullstack{
		SpecsDir: t.TempDir(),
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}
	err := Generate(fs, prepared.State{}, t.TempDir())
	if err == nil {
		t.Fatalf("expected copy openapi.yaml error, got nil")
	}
}

func TestGenerate_Success(t *testing.T) {
	specs := t.TempDir()
	apiDir := filepath.Join(specs, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	spec := "openapi: 3.0.0\ninfo:\n  title: t\n  version: '1'\npaths: {}\n"
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(spec), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	fs := &yongol.Fullstack{
		SpecsDir: specs,
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}
	arts := t.TempDir()
	if err := Generate(fs, prepared.State{}, arts); err != nil {
		t.Fatalf("Generate: %v", err)
	}
	mwDir := filepath.Join(arts, "backend", "internal", "middleware")
	for _, f := range []string{"openapi.yaml", "request_validator.go", "body_limit.go"} {
		if _, err := os.Stat(filepath.Join(mwDir, f)); err != nil {
			t.Errorf("expected %s written: %v", f, err)
		}
	}
}
