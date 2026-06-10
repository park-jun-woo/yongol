//ff:func feature=manifest type=test control=sequence
//ff:what Load — frontend.auth 블록이 있으면 token_field/refresh_field/refresh_op/store 파싱

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FrontendAuth(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: authed
backend:
  module: github.com/test/authed
frontend:
  framework: react
  auth:
    token_field: access_token
    refresh_field: refresh_token
    refresh_op: RefreshToken
    store: memory
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
	if cfg.Frontend.Auth.RefreshField != "refresh_token" {
		t.Errorf("refresh_field: want refresh_token, got %q", cfg.Frontend.Auth.RefreshField)
	}
	if cfg.Frontend.Auth.RefreshOp != "RefreshToken" {
		t.Errorf("refresh_op: want RefreshToken, got %q", cfg.Frontend.Auth.RefreshOp)
	}
	if cfg.Frontend.Auth.Store != "memory" {
		t.Errorf("store: want memory, got %q", cfg.Frontend.Auth.Store)
	}
	if got := cfg.Frontend.Auth.ResolvedStore(); got != "memory" {
		t.Errorf("ResolvedStore: want memory, got %q", got)
	}
}
