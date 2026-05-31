//ff:func feature=agent type=test control=sequence
//ff:what TestFixFile — read 에러 / 전체파일 수정 성공·LLM에러·빈응답 / split 레이어 위임 분기 검증
package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

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
