//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockJWTSecret — JWT secret 환경변수 읽기 블록

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockJWTSecret_DefaultEnv(t *testing.T) {
	block := blockJWTSecret(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}})
	if block.Name != "jwt-secret" {
		t.Errorf("name = %q, want jwt-secret", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `os.Getenv("JWT_SECRET")`) {
		t.Errorf("default env should be JWT_SECRET, got:\n%s", body)
	}
	if !strings.Contains(body, "len(v) < 32") {
		t.Errorf("must enforce minimum length 32, got:\n%s", body)
	}
}

func TestBlockJWTSecret_CustomEnv(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Auth: &pmanifest.Auth{SecretEnv: "APP_JWT"}},
	}}
	block := blockJWTSecret(fs)
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `os.Getenv("APP_JWT")`) {
		t.Errorf("custom env should be APP_JWT, got:\n%s", body)
	}
}
