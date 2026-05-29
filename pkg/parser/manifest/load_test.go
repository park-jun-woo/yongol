//ff:func feature=manifest type=parser control=sequence
//ff:what manifest.yaml 정상 로드 및 claims 파싱 검증
package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: testapp
backend:
  module: github.com/test/testapp
  auth:
    type: jwt
    claims:
      UserID: "user_id:int64"
      Role: "role"
frontend:
  framework: react
`
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644)

	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}

	if cfg.APIVersion != "yongol/v1" {
		t.Errorf("APIVersion = %q, want %q", cfg.APIVersion, "yongol/v1")
	}
	if cfg.Metadata.Name != "testapp" {
		t.Errorf("Metadata.Name = %q, want %q", cfg.Metadata.Name, "testapp")
	}
	if cfg.Backend.Module != "github.com/test/testapp" {
		t.Errorf("Backend.Module = %q", cfg.Backend.Module)
	}
	if cfg.Backend.Auth == nil {
		t.Fatal("Backend.Auth is nil")
	}
	if len(cfg.Backend.Auth.Claims) != 2 {
		t.Errorf("Claims count = %d, want 2", len(cfg.Backend.Auth.Claims))
	}
	if def, ok := cfg.Backend.Auth.Claims["UserID"]; !ok || def.Key != "user_id" || def.GoType != "int64" {
		t.Errorf("Claims[UserID] = %+v", def)
	}
}
