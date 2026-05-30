//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestRetryFailedOp — feat 미존재 early return / feat 존재+relativeLines+LLM 에러 분기 검증

package agent

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestRetryFailedOpUnknownOp(t *testing.T) {
	// An op missing from featByOp returns immediately without output.
	var out bytes.Buffer
	retryFailedOp("Unknown", map[string]features.Feature{}, nil, errors.New("verify"),
		map[string]any{}, map[string][]string{}, map[string]string{}, Config{}, &out)
	if out.Len() != 0 {
		t.Errorf("expected no output for unknown op, got: %q", out.String())
	}
}

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

func TestRetryFailedOpNoRelativeLines(t *testing.T) {
	// nil relativeLines skips the lookup (rl stays -1) but still proceeds to the
	// LLM call, which fails for the unsupported backend.
	featByOp := map[string]features.Feature{
		"CreateUser": {Op: "CreateUser", Path: "/users", Table: "users"},
	}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m", SpecsDir: t.TempDir()}
	retryFailedOp("CreateUser", featByOp, nil, errors.New("boom"),
		map[string]any{}, map[string][]string{}, map[string]string{}, cfg, &out)
	if !strings.Contains(out.String(), "LLM error") {
		t.Errorf("expected LLM-error message, got: %q", out.String())
	}
}
