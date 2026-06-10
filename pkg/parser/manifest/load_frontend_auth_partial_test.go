//ff:func feature=manifest type=test control=sequence
//ff:what Load — frontend.auth 에 token_field 만 선언 시 나머지는 빈 값 + store 기본값

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FrontendAuthPartial(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: partial
backend:
  module: github.com/test/partial
frontend:
  framework: react
  auth:
    token_field: access_token
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Frontend.Auth == nil {
		t.Fatalf("expected non-nil Auth")
	}
	if cfg.Frontend.Auth.TokenField != "access_token" {
		t.Errorf("token_field: want access_token, got %q", cfg.Frontend.Auth.TokenField)
	}
	if cfg.Frontend.Auth.RefreshField != "" {
		t.Errorf("refresh_field: want empty, got %q", cfg.Frontend.Auth.RefreshField)
	}
	if cfg.Frontend.Auth.RefreshOp != "" {
		t.Errorf("refresh_op: want empty, got %q", cfg.Frontend.Auth.RefreshOp)
	}
	if cfg.Frontend.Auth.Store != "" {
		t.Errorf("store: want empty, got %q", cfg.Frontend.Auth.Store)
	}
	if got := cfg.Frontend.Auth.ResolvedStore(); got != "localStorage" {
		t.Errorf("ResolvedStore default: want localStorage, got %q", got)
	}
}
