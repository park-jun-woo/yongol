//ff:func feature=agent type=test control=sequence
//ff:what TestLoadAPIKey — 환경변수 우선, XDG credentials.yaml fallback, 미존재 시 에러 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAPIKeyHomeFallback(t *testing.T) {
	// Empty XDG_CONFIG_HOME forces the UserHomeDir branch: the file is read from
	// $HOME/.config/yongol/credentials.yaml.
	home := t.TempDir()
	t.Setenv("XAI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("HOME", home)
	ydir := filepath.Join(home, ".config", "yongol")
	if err := os.MkdirAll(ydir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(ydir, "credentials.yaml"), []byte("xai: home-key\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := loadAPIKey("xai")
	if err != nil || got != "home-key" {
		t.Fatalf("home fallback = %q, %v; want home-key", got, err)
	}
}
