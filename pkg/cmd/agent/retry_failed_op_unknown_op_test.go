//ff:func feature=agent type=test control=sequence
//ff:what TestRetryFailedOp — feat 미존재 early return / feat 존재+relativeLines+LLM 에러 분기 검증
package agent

import (
	"bytes"
	"errors"
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
