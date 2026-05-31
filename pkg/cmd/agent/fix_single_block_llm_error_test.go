//ff:func feature=agent type=test control=sequence
//ff:what TestFixSingleBlock — default 레이어 / extract 실패→generateNewBlock / 블록추출 후 LLM 에러 분기 검증
package agent

import (
	"bytes"
	"strings"
	"testing"
)

func TestFixSingleBlockLLMError(t *testing.T) {
	// A valid OpenAPI block is extracted, then llmCall fails (unsupported
	// backend), exercising the LLM-error branch.
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n      responses:\n        '200': {}\n"
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "none"}
	got := fixSingleBlock(&out, cfg, layerOpenAPI, "api/openapi.yaml", "/tmp/api/openapi.yaml",
		&content, "ListUsers", "desc", "path", []string{"S-01: error"})
	if got {
		t.Fatal("expected false when LLM call fails")
	}
	if !strings.Contains(out.String(), "skipped block") {
		t.Errorf("expected skipped-block message, got: %q", out.String())
	}
}
