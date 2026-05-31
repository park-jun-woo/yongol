//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRunGoModTidySuccess — 최소 유효 go.mod 에서 runGoModTidy 성공 확인
package gogin

import (
	"os"
	"path/filepath"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateGoMod_MkdirError(t *testing.T) {
	dir := t.TempDir()
	// artifactsDir is a regular file -> MkdirAll(artifactsDir/backend) fails.
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}
	err := generateGoMod(fs, "example.com/app", blocker)
	if err == nil {
		t.Fatalf("expected mkdir error when backend parent is a file, got nil")
	}
}
