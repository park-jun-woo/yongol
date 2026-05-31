//ff:func feature=agent type=test control=sequence
//ff:what TestRetryFailedOp — feat 미존재 early return / feat 존재+relativeLines+LLM 에러 분기 검증
package agent

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestRetryFailedOpLLMError(t *testing.T) {
	// A known op with a relativeLines hint reaches llmCallWithNumCtx, which fails
	// for an unsupported backend, exercising the relativeLines lookup and the
	// LLM-error branch.
	featByOp := map[string]features.Feature{
		"CreateUser": {Op: "CreateUser", Path: "/users", Table: "users"},
	}
	relativeLines := map[string]int{"CreateUser": 3}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m", SpecsDir: t.TempDir()}
	retryFailedOp("CreateUser", featByOp, relativeLines, errors.New("boom"),
		map[string]any{}, map[string][]string{}, map[string]string{}, cfg, &out)
	if !strings.Contains(out.String(), "LLM error") {
		t.Errorf("expected LLM-error message, got: %q", out.String())
	}
}
