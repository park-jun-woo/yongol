//ff:func feature=agent type=test control=selection
//ff:what TestGenerateNewBlock — SSaC 부재 skip / SSaC 존재+LLM 에러 skip 분기 검증

package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateNewBlockSSaCNotFound(t *testing.T) {
	// No SSaC file under <specsDir>/service/*/ → early skip, returns false.
	dir := t.TempDir()
	abs := filepath.Join(dir, "api", "openapi.yaml")
	content := "paths: {}\n"
	var out bytes.Buffer
	if generateNewBlock(&out, Config{}, layerOpenAPI, "api/openapi.yaml", abs, &content, "ListUsers", "d", "p") {
		t.Fatal("expected false when SSaC file is missing")
	}
	if !strings.Contains(out.String(), "SSaC file not found") {
		t.Errorf("expected SSaC-not-found message, got: %q", out.String())
	}
}

func TestGenerateNewBlockLLMError(t *testing.T) {
	// SSaC file present but an unsupported backend makes llmCall fail.
	dir := t.TempDir()
	svc := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(svc, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(svc, "ListUsers.ssac"), []byte("func ListUsers() {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	abs := filepath.Join(dir, "api", "openapi.yaml")
	content := "paths: {}\n"
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "none"}
	if generateNewBlock(&out, cfg, layerOpenAPI, "api/openapi.yaml", abs, &content, "ListUsers", "d", "p") {
		t.Fatal("expected false when LLM call fails")
	}
	if !strings.Contains(out.String(), "LLM error") {
		t.Errorf("expected LLM-error message, got: %q", out.String())
	}
}
