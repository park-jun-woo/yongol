//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestGenerate_DomainMode — 도메인 모드 per-domain validator/embed 방출 + 단일 사이트 파일 미방출 (BUG-142)

package middleware

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestGenerate_DomainMode(t *testing.T) {
	fs := domainedValidatorFS(t)
	arts := t.TempDir()
	if err := Generate(fs, prepared.State{}, arts); err != nil {
		t.Fatalf("Generate: %v", err)
	}
	mwDir := filepath.Join(arts, "backend", "internal", "middleware")
	// Per-domain artifacts present.
	for _, f := range []string{
		"openapi_public.yaml", "request_validator_public.go",
		"openapi_admin.yaml", "request_validator_admin.go",
	} {
		if _, err := os.Stat(filepath.Join(mwDir, f)); err != nil {
			t.Errorf("expected %s written: %v", f, err)
		}
	}
	// Single-site artifacts MUST NOT be emitted in domain mode.
	for _, f := range []string{"openapi.yaml", "request_validator.go"} {
		if _, err := os.Stat(filepath.Join(mwDir, f)); !os.IsNotExist(err) {
			t.Errorf("domain mode must not emit %s (err=%v)", f, err)
		}
	}
	src, err := os.ReadFile(filepath.Join(mwDir, "request_validator_public.go"))
	if err != nil {
		t.Fatalf("read public validator: %v", err)
	}
	body := string(src)
	for _, must := range []string{
		"func RequestValidatorPublic() (gin.HandlerFunc, error)",
		"//go:embed openapi_public.yaml",
		"var openapiSpecPublic []byte",
		"loader.LoadFromData(openapiSpecPublic)",
		`bypassPrefixes := []string{"/health", "/ready", "/metrics"}`,
	} {
		if !strings.Contains(body, must) {
			t.Errorf("missing %q in:\n%s", must, body)
		}
	}
}
