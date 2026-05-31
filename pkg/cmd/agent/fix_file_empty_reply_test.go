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
