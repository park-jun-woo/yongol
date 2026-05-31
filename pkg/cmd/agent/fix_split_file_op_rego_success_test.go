//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixSplitFileOpRegoSuccess(t *testing.T) {
	content := "package authz\n\nallow if {\n  input.action == \"ListUsers\"\n}\n"
	fixed := "allow if {\n  input.action == \"ListUsers\"\n  input.role == \"admin\"\n}"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM(fixed, nil), Config{}, layerRego)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(content, "input.role") {
		t.Errorf("rego content not merged: %q", content)
	}
}
