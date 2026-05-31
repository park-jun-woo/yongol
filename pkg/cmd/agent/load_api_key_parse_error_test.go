//ff:func feature=agent type=test control=sequence
//ff:what TestLoadAPIKey — 환경변수 우선, XDG credentials.yaml fallback, 미존재 시 에러 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAPIKeyParseError(t *testing.T) {
	// A malformed credentials.yaml triggers the yaml.Unmarshal error branch.
	cfg := t.TempDir()
	t.Setenv("XAI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", cfg)
	ydir := filepath.Join(cfg, "yongol")
	if err := os.MkdirAll(ydir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Sequence node where a map is expected → unmarshal into map[string]string fails.
	if err := os.WriteFile(filepath.Join(ydir, "credentials.yaml"), []byte("- a\n- b\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadAPIKey("xai"); err == nil {
		t.Error("expected parse error for malformed credentials.yaml")
	}
}
