//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldRegoBatch — 미지원 backend LLM 호출 실패 시 배치 인덱스 포함 에러 전파 검증

package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldRegoBatch(t *testing.T) {
	batch := []features.Feature{{Op: "ListUsers", Path: "/users", Table: "users"}}
	cfg := Config{Backend: "unsupported-backend", Model: "m"}

	out, err := scaffoldRegoBatch(batch, 3, "system prompt", cfg)
	if err == nil {
		t.Fatal("expected error from LLM call with unsupported backend")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
	if !strings.Contains(err.Error(), "scaffold rego batch 3") {
		t.Errorf("error should include batch index: %v", err)
	}
}
