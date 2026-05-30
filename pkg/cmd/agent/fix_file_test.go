//ff:func feature=agent type=test control=sequence
//ff:what TestFixFile — read 에러 / 전체파일 수정 성공·LLM에러·빈응답 / split 레이어 위임 분기 검증

package agent

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixFileReadError(t *testing.T) {
	// A non-existent target file makes os.ReadFile fail.
	err := fixFile(t.TempDir(), &features.FeaturesFile{}, "missing.ssac", nil,
		func(b, m, s, u string) (string, error) { return "", nil }, Config{})
	if err == nil {
		t.Fatal("expected read error for missing file")
	}
	if !strings.Contains(err.Error(), "read") {
		t.Errorf("expected read error, got: %v", err)
	}
}

func TestFixFileWholeFileSuccess(t *testing.T) {
	dir := t.TempDir()
	svc := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(svc, 0o755); err != nil {
		t.Fatal(err)
	}
	rel := filepath.Join("service", "auth", "Login.ssac")
	if err := os.WriteFile(filepath.Join(dir, rel), []byte("old content"), 0o644); err != nil {
		t.Fatal(err)
	}

	llm := func(b, m, s, u string) (string, error) { return "```\nnew content\n```", nil }
	err := fixFile(dir, &features.FeaturesFile{}, rel, nil, llm, Config{Backend: "x", Model: "y"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, _ := os.ReadFile(filepath.Join(dir, rel))
	if !strings.Contains(string(got), "new content") {
		t.Errorf("file not updated, got: %q", got)
	}
}

func TestFixFileLLMError(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata: {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	llm := func(b, m, s, u string) (string, error) { return "", errors.New("boom") }
	err := fixFile(dir, &features.FeaturesFile{}, "manifest.yaml", nil, llm, Config{})
	if err == nil {
		t.Fatal("expected LLM error")
	}
	if !strings.Contains(err.Error(), "LLM") {
		t.Errorf("expected LLM error, got: %v", err)
	}
}

func TestFixFileEmptyReply(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata: {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	// A reply that strips down to empty triggers the empty-response branch.
	llm := func(b, m, s, u string) (string, error) { return "```\n```", nil }
	err := fixFile(dir, &features.FeaturesFile{}, "manifest.yaml", nil, llm, Config{})
	if err == nil {
		t.Fatal("expected empty-response error")
	}
	if !strings.Contains(err.Error(), "empty") {
		t.Errorf("expected empty-response error, got: %v", err)
	}
}

func TestFixFileSplitDelegation(t *testing.T) {
	// An OpenAPI file routes to fixSplitFile. With diagnostics carrying no
	// operationId, fixSplitFile returns its "no operationId" error, proving the
	// split-layer delegation branch was taken.
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte("paths: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	rel := filepath.Join("api", "openapi.yaml")
	diags := []diagnostic.Diagnostic{{Message: "some error", Level: diagnostic.LevelError}}
	llm := func(b, m, s, u string) (string, error) { return "x", nil }
	err := fixFile(dir, &features.FeaturesFile{}, rel, diags, llm, Config{})
	if err == nil {
		t.Fatal("expected delegation to fixSplitFile to surface an error")
	}
	if !strings.Contains(err.Error(), "operationId") {
		t.Errorf("expected fixSplitFile no-operationId error, got: %v", err)
	}
}
