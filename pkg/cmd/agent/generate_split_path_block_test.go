//ff:func feature=agent type=test control=sequence
//ff:what TestGenerateSplitPathBlock — 미지원 backend 로 첫 parameters LLM 호출 실패 시 에러 전파 검증

package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// TestGenerateSplitPathBlock drives the first failure path: with an unsupported
// backend, the initial callStepWithRetry (parameters) fails and the error is
// wrapped with the "parameters:" prefix. The subsequent requestBody/schema200
// branches require a live LLM endpoint and are not deterministically reachable.
func TestGenerateSplitPathBlock(t *testing.T) {
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	_, err := generateSplitPathBlock(features.Feature{Op: "CreateUser", Path: "/users"}, "", cfg)
	if err == nil {
		t.Fatal("expected error when backend is unsupported")
	}
	if !strings.Contains(err.Error(), "parameters") {
		t.Errorf("expected parameters error prefix, got: %v", err)
	}
}
