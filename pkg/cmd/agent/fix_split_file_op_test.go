//ff:func feature=agent type=test control=selection
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증

package agent

import (
	"errors"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func mockLLM(reply string, err error) LLMCallFunc {
	return func(b, m, s, u string) (string, error) { return reply, err }
}

func TestFixSplitFileOpExtractError(t *testing.T) {
	content := "paths: {}\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "Missing", nil, nil, nil,
		mockLLM("x", nil), Config{}, layerOpenAPI)
	if err == nil {
		t.Fatal("expected extract error for missing op")
	}
}

func TestFixSplitFileOpLLMError(t *testing.T) {
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM("", errors.New("boom")), Config{}, layerOpenAPI)
	if err == nil || !strings.Contains(err.Error(), "boom") {
		t.Fatalf("expected LLM error, got: %v", err)
	}
}

func TestFixSplitFileOpEmptyReply(t *testing.T) {
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM("```\n```", nil), Config{}, layerOpenAPI)
	if err == nil || !strings.Contains(err.Error(), "empty") {
		t.Fatalf("expected empty-response error, got: %v", err)
	}
}

func TestFixSplitFileOpMergeError(t *testing.T) {
	// A fixed block lacking operationId fails the OpenAPI merge validation.
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM("/users:\n  get:\n    summary: no opid\n", nil), Config{}, layerOpenAPI)
	if err == nil {
		t.Fatal("expected merge error for block without operationId")
	}
}

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

func TestFixSplitFileOpHurlSuccess(t *testing.T) {
	content := "# ListUsers\nGET https://example.com/users\nHTTP 200\n"
	fixed := "# ListUsers\nGET https://example.com/users\nHTTP 200\n# extra"
	err := fixSplitFileOp(t.TempDir(), &features.FeaturesFile{}, &content, "ListUsers", nil, nil, nil,
		mockLLM(fixed, nil), Config{}, layerHurl)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
