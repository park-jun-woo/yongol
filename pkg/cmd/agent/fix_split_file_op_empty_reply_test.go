//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixSplitFileOpEmptyReply(t *testing.T) {
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM("```\n```", nil), Config{}, layerOpenAPI)
	if err == nil || !strings.Contains(err.Error(), "empty") {
		t.Fatalf("expected empty-response error, got: %v", err)
	}
}
