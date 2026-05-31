//ff:func feature=agent type=test control=sequence
//ff:what TestFixFile — read 에러 / 전체파일 수정 성공·LLM에러·빈응답 / split 레이어 위임 분기 검증
package agent

import (
	"strings"
	"testing"

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
