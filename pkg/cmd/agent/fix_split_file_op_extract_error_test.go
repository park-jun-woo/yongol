//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixSplitFileOpExtractError(t *testing.T) {
	content := "paths: {}\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "Missing", nil, nil, nil,
		mockLLM("x", nil), Config{}, layerOpenAPI)
	if err == nil {
		t.Fatal("expected extract error for missing op")
	}
}
