//ff:func feature=gen-filefunc type=test control=iteration dimension=1
//ff:what TestBuildFeatureEntries — SSOT+internal 디렉토리 병합 feature 맵 생성 검증
package filefunc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildFeatureEntries(t *testing.T) {
	backendDir := t.TempDir()
	internalDir := filepath.Join(backendDir, "internal")
	// internal/ has a "store" subdirectory (feature from disk) and a stray file.
	if err := os.MkdirAll(filepath.Join(internalDir, "store"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(internalDir, "readme.txt"), []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{Feature: "workflow"},
		},
		ProjectFuncSpecs: []funcspec.FuncSpec{
			{Package: "auth", Description: "authentication"},
		},
	}

	out := buildFeatureEntries(fs, backendDir)

	// SSOT-derived features present.
	if _, ok := out["workflow"]; !ok {
		t.Errorf("expected SSOT feature 'workflow' present: %v", out)
	}
	if got := out["auth"]; got != "authentication" {
		t.Errorf("expected funcspec description for auth, got %q", got)
	}
	// internal/ subdir merged.
	if _, ok := out["store"]; !ok {
		t.Errorf("expected internal dir 'store' present: %v", out)
	}
	// anchor entries always present.
	for _, anchor := range []string{"gen-filefunc", "runtime-middleware", "main"} {
		if _, ok := out[anchor]; !ok {
			t.Errorf("expected anchor feature %q present: %v", anchor, out)
		}
	}
}
