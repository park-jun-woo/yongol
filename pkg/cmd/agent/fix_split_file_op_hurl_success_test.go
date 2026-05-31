//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixSplitFileOpHurlSuccess(t *testing.T) {
	content := "# ListUsers\nGET https://example.com/users\nHTTP 200\n"
	fixed := "# ListUsers\nGET https://example.com/users\nHTTP 200\n# extra"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM(fixed, nil), Config{}, layerHurl)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
