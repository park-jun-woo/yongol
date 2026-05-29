//ff:func feature=agent type=test control=sequence
//ff:what TestLoadAPIKey — 환경변수 우선, XDG credentials.yaml fallback, 미존재 시 에러 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAPIKey(t *testing.T) {
	// Priority 1: environment variable wins.
	t.Setenv("XAI_API_KEY", "env-key")
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	if got, err := loadAPIKey("xai"); err != nil || got != "env-key" {
		t.Fatalf("env path = %q, %v; want env-key", got, err)
	}

	// Priority 2: credentials.yaml under XDG config dir.
	cfg := t.TempDir()
	t.Setenv("XAI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", cfg)
	ydir := filepath.Join(cfg, "yongol")
	if err := os.MkdirAll(ydir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(ydir, "credentials.yaml"), []byte("xai: file-key\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got, err := loadAPIKey("xai"); err != nil || got != "file-key" {
		t.Fatalf("file path = %q, %v; want file-key", got, err)
	}

	// Missing key entry in an otherwise present file errors.
	if _, err := loadAPIKey("gemini"); err == nil {
		t.Error("expected error when backend key absent from credentials.yaml")
	}

	// No env and no file at all errors.
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	if _, err := loadAPIKey("xai"); err == nil {
		t.Error("expected error when neither env nor credentials.yaml available")
	}
}
