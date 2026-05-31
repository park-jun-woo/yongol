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
