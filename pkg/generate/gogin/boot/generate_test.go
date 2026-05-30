//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — boot main.go + env helper 생성 success/error 경로 검증

package boot

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate(t *testing.T) {
	newFS := func() *yongol.Fullstack {
		return &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
		}
	}

	t.Run("WritesMainAndHelpers", func(t *testing.T) {
		dir := t.TempDir()
		fs := newFS()
		p := prepared.New(fs)
		if err := Generate(fs, p, dir); err != nil {
			t.Fatalf("Generate error: %v", err)
		}
		mainPath := filepath.Join(dir, "backend", "cmd", "main.go")
		info, err := os.Stat(mainPath)
		if err != nil {
			t.Fatalf("expected main.go: %v", err)
		}
		if info.Size() == 0 {
			t.Errorf("main.go is empty")
		}
	})

	t.Run("WriteMainGoError", func(t *testing.T) {
		dir := t.TempDir()
		// backend/cmd parent (backend) is a regular file -> MkdirAll fails.
		backend := filepath.Join(dir, "backend")
		if err := os.WriteFile(backend, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		fs := newFS()
		p := prepared.New(fs)
		if err := Generate(fs, p, dir); err == nil {
			t.Errorf("expected error when main.go dir cannot be created")
		}
	})

	t.Run("NilManifest", func(t *testing.T) {
		dir := t.TempDir()
		fs := &yongol.Fullstack{}
		p := prepared.New(fs)
		if err := Generate(fs, p, dir); err != nil {
			t.Fatalf("Generate with nil manifest error: %v", err)
		}
	})
}
