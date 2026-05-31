//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixSplitFileOpOpenAPISuccess(t *testing.T) {
	// desc lookup hit + filterMessagesByOp fallback + successful OpenAPI merge.
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	lookup := map[string]features.Feature{"ListUsers": {Op: "ListUsers", Desc: "list users"}}
	fixed := "  /users:\n    get:\n      operationId: ListUsers\n      summary: fixed\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers",
		[]diagnostic.Diagnostic{{Message: "err"}}, []string{"unrelated msg"}, lookup,
		mockLLM(fixed, nil), Config{}, layerOpenAPI)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(content, "summary: fixed") {
		t.Errorf("content not merged: %q", content)
	}
}
